package lab2

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type operator struct {
	Regex         string
	Arity         int
	Priority      int
	Format        string
	IsAssociative bool
}

var parseErr = errors.New("Error while parsing input string!")
var unknownOperatorErr = errors.New("Could not parse operator in the string!")

func (op operator) evaluate(token string, args []expNode) string {
	format := strings.Replace(op.Format, "%token", token, -1)
	var argStrings []interface{} = make([]interface{}, op.Arity)
	for i := 0; i < op.Arity; i++ {
		argStr := args[i].evaluate()
		isLessPrioritized := args[i].Operator.Priority < op.Priority
		isSameNonAss := args[i].Operator.Priority == op.Priority && !op.IsAssociative && i >= 1
		if isLessPrioritized || isSameNonAss {
			argStr = fmt.Sprintf("(%s)", argStr)
		}
		argStrings[i] = argStr
	}
	return fmt.Sprintf(format, argStrings...)
}

type expNode struct {
	Operator *operator
	Token    string
	Args     []expNode
}

func (node expNode) evaluate() string {
	evaled := node.Operator.evaluate(node.Token, node.Args)
	return evaled
}

var operators = []operator{
	{
		Regex:         `\+`,
		Arity:         2,
		Priority:      10,
		Format:        "%v %token %v",
		IsAssociative: true,
	},
	{
		Regex:         `\-`,
		Arity:         2,
		Priority:      10,
		Format:        "%v %token %v",
		IsAssociative: false,
	},
	{
		Regex:         `\*`,
		Arity:         2,
		Priority:      20,
		Format:        "%v %token %v",
		IsAssociative: true,
	},
	{
		Regex:         `\/`,
		Arity:         2,
		Priority:      20,
		Format:        "%v %token %v",
		IsAssociative: false,
	},

	{
		Regex:         `\^`,
		Arity:         2,
		Priority:      40,
		Format:        "%v %token %v",
		IsAssociative: false,
	},

	{
		Regex:         `sin`,
		Arity:         1,
		Priority:      50,
		Format:        "%token(%v)",
		IsAssociative: false,
	},
	{
		Regex:    `[0-9]+(\.[0-9]+)?`,
		Arity:    0,
		Priority: 100,
		Format:   "%token",
	},

	{
		Regex:    `[a-zA-Z]`,
		Arity:    0,
		Priority: 100,
		Format:   "%token",
	},
}

func parseOperator(str string) (*operator, []int, error) {
	var opLoc []int
	var operator *operator
	for _, v := range operators {
		r := regexp.MustCompile(`\A` + v.Regex)
		opLoc = r.FindStringIndex(str)
		if opLoc != nil {
			operator = &v
			break
		}
	}
	if opLoc == nil {
		return nil, nil, unknownOperatorErr
	}
	return operator, opLoc, nil
}

func parsePrefix(str string) (*expNode, string, error) {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return nil, "", parseErr
	}
	operator, opLoc, er := parseOperator(str)
	if er != nil {
		return nil, "", er
	}
	token := str[opLoc[0]:opLoc[1]]
	left := str[opLoc[1]:]

	args := make([]expNode, operator.Arity)
	for i := 0; i < operator.Arity; i++ {
		arg, leftAfterArg, err := parsePrefix(left)
		if err != nil {
			return nil, "", err
		}
		args[i] = *arg
		left = leftAfterArg
	}

	node := &expNode{
		Operator: operator,
		Token:    token,
		Args:     args,
	}
	return node, left, nil
}

func PrefixToInfix(input string) (string, error) {
	node, left, er := parsePrefix(input)
	if er != nil {
		return "", er
	}
	if len(left) > 0 {
		return "", parseErr
	}
	str := node.evaluate()
	return str, nil
}
