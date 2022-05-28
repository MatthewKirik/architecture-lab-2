package lab2

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Operator struct {
	Regex         string
	Arity         int
	Priority      int
	Format        string
	IsAssociative bool
}

var parseEr = errors.New("Error while parsing input string!")
var errWrongArity = errors.New("Arguments count of node does not match operator arity!")
var errUnknownOperator = errors.New("Could not parse operator in the string!")

func (op Operator) evaluate(token string, args []ExpNode) string {
	format := strings.Replace(op.Format, "%token", token, -1)
	var argStrings []interface{} = make([]interface{}, op.Arity)
	for i := 0; i < op.Arity; i++ {
		argStr, er := args[i].evaluate()
		if handleError(er) {
			os.Exit(-1)
		}

		isLessPrioritized := args[i].Operator.Priority < op.Priority
		isSameNonAss := args[i].Operator.Priority == op.Priority && !op.IsAssociative && i >= 1
		if isLessPrioritized || isSameNonAss {
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

func (node ExpNode) evaluate() (string, error) {
	if node.Operator.Arity != len(node.Args) {
		return "", errWrongArity
	}
	evaled := node.Operator.evaluate(node.Token, node.Args)
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
		return nil, nil, errUnknownOperator
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

func PrefixToInfix(input string) (string, error) {
	node, left, er := parsePrefix(input)
	if er != nil {
		return "", er
	}
	if len(left) > 0 {
		return "", parseEr
	}
	str, er := node.evaluate()
	if er != nil {
		return "", er
	}
	return str, nil
}
