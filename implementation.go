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
	Regex         string
	Arity         int
	Priority      int
	Format        string
	IsAssociative bool
}

const opens_group = -1
const closes_group = -2

func (op Operator) Evaluate(token string, args []ExpNode) string {
	format := strings.Replace(op.Format, "%token", token, -1)
	var argStrings []interface{} = make([]interface{}, op.Arity)
	for i := 0; i < op.Arity; i++ {
		argStr, er := args[i].Evaluate()
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
		Format:        "%v%token%v",
		IsAssociative: false,
	},

	{
		Regex:         `\-`,
		Arity:         1,
		Priority:      100,
		Format:        "%token%v",
		IsAssociative: false,
	},

	{
		Regex:         `\(`,
		Arity:         opens_group,
		Priority:      200,
		Format:        "",
		IsAssociative: false,
	},
	{
		Regex:         `\)`,
		Arity:         closes_group,
		Priority:      200,
		Format:        "",
		IsAssociative: false,
	},

	{
		Regex:    `[0-9]+(\.[0-9]+)?`,
		Arity:    0,
		Priority: 100,
		Format:   "%token",
	},
}

var ErrUnknownOperator = errors.New("Could not parse operator in the string!")

type OperatorLocation struct {
	Operator Operator
	Location []int
}

func parseOperators(str string) ([]OperatorLocation, error) {
	var opLocs []OperatorLocation
	for _, v := range operators {
		r := regexp.MustCompile(`\A` + v.Regex)
		opLoc := r.FindStringIndex(str)
		if len(opLoc) == 2 {
			opLocs = append(opLocs, OperatorLocation{
				Operator: v,
				Location: opLoc,
			})
		}
	}
	if len(opLocs) == 0 {
		return nil, ErrUnknownOperator
	}
	return opLocs, nil
}

func parsePrefixArgs(str string, arity int) ([]ExpNode, string, error) {
	var args []ExpNode
	var left string = str
	nextArgGreedy := false
	for i := 0; i < arity; i++ {
		arg, leftAfterArg, err := parsePrefix(left, nextArgGreedy)
		if err != nil {
			return nil, "", err
		}
		left = leftAfterArg
		nextArgGreedy = false

		if arg.Operator.Arity == opens_group {
			i--
			nextArgGreedy = true
			continue
		}
		if arg.Operator.Arity == closes_group {
			break
		}
		args = append(args, *arg)
	}
	return args, left, nil
}

var ErrArgAmountMismatch = errors.New("Arguments amount mismatch!")

func parsePrefix(str string, greedy bool) (*ExpNode, string, error) {
	str = strings.TrimSpace(str)
	ops, er := parseOperators(str)
	if er != nil {
		return nil, "", er
	}
	opLoc := ops[0]
	var arity int
	if !greedy {
		for i := 1; i < len(ops)-1; i++ {
			if ops[i].Operator.Arity > opLoc.Operator.Arity {
				opLoc = ops[i]
			}
		}
		arity = opLoc.Operator.Arity
	} else {
		arity = math.MaxInt32
	}

	token := str[opLoc.Location[0]:opLoc.Location[1]]
	args, left, er := parsePrefixArgs(str[opLoc.Location[1]:], arity)
	if er != nil {
		return nil, "", er
	}
	if greedy {
		argsCount := len(args)
		found := false
		for _, v := range ops {
			if v.Operator.Arity == argsCount {
				opLoc = v
				found = true
				break
			}
		}
		if !found {
			return nil, "", er
		}
	}

	node := &ExpNode{
		Operator: &opLoc.Operator,
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
	fmt.Println(str)
	return "TODO", fmt.Errorf("TODO")
}
