package main

import (
	"flag"

	lab2 "github.com/roman-mazur/architecture-lab-2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	// TODO: Add other flags support for input and output configuration.
)

func main() {
	flag.Parse()
	lab2.PrefixToInfix("/ * / 22 12 44 * 1 + 10 1")         // 22 / 12 * 44 / (1 * (10 + 1))
	lab2.PrefixToInfix("/ / / 12 12 12 12")                 // 12 / 12 / 12 / 12
	lab2.PrefixToInfix("+ / 99 + 51 1 12")                  // 99 / (51 + 1) + 12
	lab2.PrefixToInfix("+ + ^ 12 1 / 122 1 * * 11 2 - 1 0") // 12^1 + 122 / 1 + 11 * 2 * (1 - 0)
	lab2.PrefixToInfix("^ 55 + - + * 12 12 / 12 12 1 2")    // 55 ^ (12 * 12 + 12 / 12 - 1 + 2)
	lab2.PrefixToInfix("- - + 12 111 + 12 1 - 10 2")        // ((12+111)) - (12 + 1) - (10 - 2)
}
