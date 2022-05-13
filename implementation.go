package lab2

import (
	"errors"
	"fmt"
	"math"
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

const opens_group = -1
const closes_group = -2

func (op Operator) Evaluate(token string, args []ExpNode) string {
	format := strings.Replace(op.Format, "%token", token, -1)
	var arity int
	if op.Arity == opens_group {
		arity = 1
	} else {
		arity = op.Arity
	}
	var argStrings []interface{}
	for i := 0; i < arity; i++ {
		argStr, er := args[i].Evaluate()
		if handleError(er) {
			os.Exit(-1)
		}
		if args[i].Operator.Priority < op.Priority {
			argStr = fmt.Sprintf("(%s)", argStr)
		}
		argStrings = append(argStrings, argStr)
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
	if node.Operator.Arity > 0 && node.Operator.Arity != len(node.Args) {
		return "", ErrWrongArity
	}
	fmt.Println(node.Token)
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
		Regex:    `\-`,
		Arity:    1,
		Priority: 30,
		Format:   "%token%v",
	},
	{
		Regex:    `\+`,
		Arity:    1,
		Priority: 30,
		Format:   "%v",
	},

	{
		Regex:    `\^`,
		Arity:    1,
		Priority: 40,
		Format:   "%token%v",
	},

	{
		Regex:    `\(`,
		Arity:    opens_group,
		Priority: 50,
		Format:   "%v",
	},
	{
		Regex:    `\)`,
		Arity:    closes_group,
		Priority: 50,
		Format:   "",
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
		fmt.Print(len(str))
		return nil, nil, ErrUnknownOperator
	}
	return operator, opLoc, nil
}

func parsePrefix(str string, greedy bool) (*ExpNode, string, error) {
	str = strings.TrimSpace(str)
	operator, opLoc, er := parseOperator(str)
	if er != nil {
		return nil, "", er
	}
	token := str[opLoc[0]:opLoc[1]]
	left := str[opLoc[1]:]

	var args []ExpNode
	var arity int
	argGreedy := false
	if greedy {
		arity = math.MaxInt32
	} else if operator.Arity == opens_group {
		arity = 1
		argGreedy = true
	} else if operator.Arity == closes_group {
		arity = 0
	} else {
		arity = operator.Arity
	}
	fmt.Println("Entering " + operator.Regex)
	for i := 0; i < arity; i++ {
		arg, leftAfterArg, err := parsePrefix(left, argGreedy)
		if err != nil {
			return nil, "", err
		}
		left = leftAfterArg
		if arg.Operator.Arity == closes_group {
			break
		}
		args = append(args, *arg)
	}
	fmt.Println(len(args))
	fmt.Println("Exiting " + operator.Regex)

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
	node, _, _ := parsePrefix(input, false)
	str, _ := node.Evaluate()
	fmt.Print(str)
	return "TODO", fmt.Errorf("TODO")
}
