package main

import (
	"flag"
	"strings"
)

var (
	inputExpression = flag.String("e", "xx", "Prefix expression that should be converted")
	inputFile       = flag.String("f", "", "Input file with prefix expression")
	outputFile      = flag.String("o", "", "Result file with infix expression")
)

func main() {
	flag.Parse()

	*inputExpression = strings.TrimSpace(*inputExpression)
	*inputFile = strings.TrimSpace(*inputFile)
	*outputFile = strings.TrimSpace(*outputFile)

	// inputIsNotSpecified = len(*inputFile) == 0 &&
	// 	len(*inputExpression) == 0
	// if inputIsNotSpecified {
	// 	// TODO: print error
	// }

	// manyInputsSpecified := len(*inputFile) > 0 &&
	// 	len(*inputExpression) > 0
	// if manyInputsSpecified {
	// 	// TODO: print error
	// }

	if *inputExpression != "" && *inputFile != "" {
		// TODO: print error many inputs specified

	} else if *inputExpression != "" && *inputFile == "" {
		// TODO create new reader

	} else if *inputExpression == "" && *inputFile != "" {
		// try open input file, handle error

	} else {
		// TODO: Print error no input was specified

	}

	if *outputFile != "" {
		// Create output file

	} else {
		// Set defaut output as STDOUT

	}

	// fmt.Printf("value = \"%s\"\n", *inputExpression)
	// fmt.Printf("value = \"%s\"\n", *inputFile)
	// fmt.Printf("value = \"%s\"\n", *outputFile)

	// prints to stderr
	// fmt.Fprintf(os.Stderr, "number of foo: %d", 1)
	// fmt.Fprintln(os.Stderr, "hello world")

}
