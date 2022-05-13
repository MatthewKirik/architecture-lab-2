package lab2

import (
	"errors"
	"fmt"
	"os"
)

type Operator struct {
	Symbol   rune
	Arity    int
	Priority int
	Format   string
}

func (op Operator) Evaluate(args []ExpNode) string {
	var argStrings []interface{} = make([]interface{}, op.Arity)
	for i := 0; i < op.Arity; i++ {
		argStr, er := args[i].Evaluate()
		if handleError(er) {
			os.Exit(-1)
		}
		argStrings[i] = argStr
	}
	return fmt.Sprintf(op.Format, argStrings...)
}

type ExpNode struct {
	Operator Operator
	Token    string
	Args     []ExpNode
}

var ErrWrongArity = errors.New("Arguments count of node does not match operator arity!")

func (node ExpNode) Evaluate() (string, error) {
	if node.Operator.Arity != len(node.Args) {
		return "", ErrWrongArity
	}
	return node.Operator.Evaluate(node.Args), nil
}

func handleError(err error) bool {
	if err != nil {
		fmt.Errorf("An error occured:  %v", err)
		return true
	}
	return false
}

var operators = map[string]Operator{
	"+": {
		Symbol:   '+',
		Arity:    2,
		Priority: 10,
		Format:   "%v + %v",
	},
	"-": {
		Symbol:   '-',
		Arity:    2,
		Priority: 10,
		Format:   "%v - %v",
	},
	"*": {
		Symbol:   '*',
		Arity:    2,
		Priority: 20,
		Format:   "%v * %v",
	},
	"/": {
		Symbol:   '/',
		Arity:    2,
		Priority: 20,
		Format:   "%v / %v",
	},
}

// TODO: document this function.
// PrefixToInfix converts
func PrefixToInfix(input string) (string, error) {
	return "TODO", fmt.Errorf("TODO")
}
