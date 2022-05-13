package lab2

import (
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

// TODO: document this function.
// PrefixToInfix converts
func PrefixToInfix(input string) (string, error) {
	return "TODO", fmt.Errorf("TODO")
}
