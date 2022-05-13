package lab2

import (
	"errors"
	"fmt"
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

// TODO: document this function.
// PrefixToInfix converts
func PrefixToInfix(input string) (string, error) {
	return "TODO", fmt.Errorf("TODO")
}
