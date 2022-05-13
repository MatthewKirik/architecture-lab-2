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
	Evaluate func(int, ...ExpNode) string
}

func DefaultEvaluate(arity int, format string, args []ExpNode) string {
	var argStrings []interface{} = make([]interface{}, arity)
	for i := 0; i < arity; i++ {
		argStr, er := args[i].Evaluate()
		if handleError(er) {
			os.Exit(-1)
		}
		argStrings[i] = argStr
	}
	return fmt.Sprintf("%s + %s", argStrings...)
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
	return node.Operator.Evaluate(node.Operator.Arity, node.Args...), nil
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
		Evaluate: func(arity int, args ...ExpNode) string {
			return DefaultEvaluate(arity, "%v + %v", args)
		},
	},
	"-": {
		Symbol:   '-',
		Arity:    2,
		Priority: 10,
		Evaluate: func(arity int, args ...ExpNode) string {
			return DefaultEvaluate(arity, "%v - %v", args)
		},
	},
	"*": {
		Symbol:   '*',
		Arity:    2,
		Priority: 20,
		Evaluate: func(arity int, args ...ExpNode) string {
			return DefaultEvaluate(arity, "%v * %v", args)
		},
	},
	"/": {
		Symbol:   '/',
		Arity:    2,
		Priority: 20,
		Evaluate: func(arity int, args ...ExpNode) string {
			return DefaultEvaluate(arity, "%v / %v", args)
		},
	},
}

// TODO: document this function.
// PrefixToInfix converts
func PrefixToInfix(input string) (string, error) {
	return "TODO", fmt.Errorf("TODO")
}
