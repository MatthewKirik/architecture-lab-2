package main

import (
	"flag"
	"io"
	"log"
	"os"
	"strings"

	lab2 "github.com/MatthewKirik/architecture-lab-2"
)

var (
	inputExpression = flag.String("e", "", "Prefix expression that should be converted")
	inputFilepath   = flag.String("f", "", "Input file with prefix expression")
	outputFilepath  = flag.String("o", "", "Result file with infix expression")
)

func main() {
	flag.Parse()

	*inputExpression = strings.TrimSpace(*inputExpression)
	*inputFilepath = strings.TrimSpace(*inputFilepath)
	*outputFilepath = strings.TrimSpace(*outputFilepath)

	var (
		reader = io.Reader(nil)
		writer = io.Writer(nil)
		err    = error(nil)
	)

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

	if *inputExpression != "" && *inputFilepath != "" {
		log.Fatal("You have specified too many inputs")
	} else if *inputExpression != "" && *inputFilepath == "" {
		reader = strings.NewReader(*inputExpression)
	} else if *inputExpression == "" && *inputFilepath != "" {
		reader, err = os.Open(*inputFilepath)
		if err != nil {
			log.Fatalf("Cannot open file with path: '%s'", *inputFilepath)
		}
	} else {
		log.Fatal("You have not specified any input")
	}

	if *outputFilepath != "" {
		writer, err = os.Create(*outputFilepath)
		if err != nil {
			log.Fatalf("Cannot create file with path: '%s'", *outputFilepath)
		}
	} else {
		writer = os.Stdout
	}

	chPtr := &lab2.ComputeHandler{
		Reader: reader,
		Writer: writer,
	}

	if errCompute := chPtr.Compute(); errCompute != nil {
		log.Fatal(errCompute)
	}

	// fmt.Printf("value = \"%s\"\n", *inputExpression)
	// fmt.Printf("value = \"%s\"\n", *inputFile)
	// fmt.Printf("value = \"%s\"\n", *outputFile)

	// prints to stderr
	// fmt.Fprintf(os.Stderr, "number of foo: %d", 1)
	// fmt.Fprintln(os.Stderr, "hello world")
}
