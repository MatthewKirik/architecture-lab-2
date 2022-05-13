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
	Evaluate func([]ExpNode) string
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
		Evaluate: func(args []ExpNode) string {
			arg1Str, er := args[0].Evaluate()
			if handleError(er) {
				os.Exit(-1)
			}
			arg2Str, er := args[1].Evaluate()
			if handleError(er) {
				os.Exit(-1)
			}
			return fmt.Sprintf("%s + %s", arg1Str, arg2Str)
		},
	},
}

// TODO: document this function.
// PrefixToInfix converts
func PrefixToInfix(input string) (string, error) {
	return "TODO", fmt.Errorf("TODO")
}
