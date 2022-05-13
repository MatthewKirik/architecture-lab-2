package lab2

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Operator struct {
	Regex    string
	Arity    int
	Priority int
	Format   string
}

func (op Operator) Evaluate(token string, args []ExpNode) string {
	format := strings.Replace(op.Format, "%token", token, -1)
	var argStrings []interface{} = make([]interface{}, op.Arity)
	for i := 0; i < op.Arity; i++ {
		argStr, er := args[i].Evaluate()
		if handleError(er) {
			os.Exit(-1)
		}
		if args[i].Operator.Priority < op.Priority {
			argStr = fmt.Sprintf("(%s)", argStr)
		}
		argStrings[i] = argStr
	}
	return fmt.Sprintf(format, argStrings...)
}

type ExpNode struct {
	Operator *Operator
	Token    string
	Args     []ExpNode
}

var ErrWrongArity = errors.New("Arguments count of node does not match operator arity!")

func (node ExpNode) Evaluate() (string, error) {
	if node.Operator.Arity != len(node.Args) {
		return "", ErrWrongArity
	}
	evaled := node.Operator.Evaluate(node.Token, node.Args)
	return evaled, nil
}

func handleError(err error) bool {
	if err != nil {
		fmt.Errorf("An error occured:  %v", err)
		return true
	}
	return false
}

var operators = []Operator{
	{
		Regex:    `\+`,
		Arity:    2,
		Priority: 10,
		Format:   "%v %token %v",
	},
	{
		Regex:    `\-`,
		Arity:    2,
		Priority: 10,
		Format:   "%v %token %v",
	},
	{
		Regex:    `\*`,
		Arity:    2,
		Priority: 20,
		Format:   "%v %token %v",
	},
	{
		Regex:    `\/`,
		Arity:    2,
		Priority: 20,
		Format:   "%v %token %v",
	},

	{
		Regex:    `\^`,
		Arity:    1,
		Priority: 40,
		Format:   "%token%v",
	},

	{
		Regex:    `[0-9]+(\.[0-9]+)?`,
		Arity:    0,
		Priority: 100,
		Format:   "%token",
	},
}

var ErrUnknownOperator = errors.New("Could not parse operator in the string!")

func parseOperator(str string) (*Operator, []int, error) {
	var opLoc []int
	var operator *Operator
	for _, v := range operators {
		r := regexp.MustCompile(`\A` + v.Regex)
		opLoc = r.FindStringIndex(str)
		if len(opLoc) == 2 {
			operator = &v
			break
		}
	}
	if len(opLoc) != 2 {
		return nil, nil, ErrUnknownOperator
	}
	return operator, opLoc, nil
}

func parsePrefix(str string) (*ExpNode, string, error) {
	str = strings.TrimSpace(str)
	operator, opLoc, er := parseOperator(str)
	if er != nil {
		return nil, "", er
	}
	token := str[opLoc[0]:opLoc[1]]
	left := str[opLoc[1]:]

	args := make([]ExpNode, operator.Arity)
	for i := 0; i < operator.Arity; i++ {
		arg, leftAfterArg, err := parsePrefix(left)
		if err != nil {
			return nil, "", err
		}
		args[i] = *arg
		left = leftAfterArg
	}

	node := &ExpNode{
		Operator: operator,
		Token:    token,
		Args:     args,
	}
	return node, left, nil
}

// TODO: document this function.
// PrefixToInfix converts
func PrefixToInfix(input string) (string, error) {
	node, _, _ := parsePrefix(input)
	str, _ := node.Evaluate()
	fmt.Print(str)
	return "TODO", fmt.Errorf("TODO")
}
