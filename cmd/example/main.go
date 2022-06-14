package main

import (
	"errors"
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

func closeFile(f *os.File) {
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func processExpression() error {
	flag.Parse()

	*inputExpression = strings.TrimSpace(*inputExpression)
	*inputFilepath = strings.TrimSpace(*inputFilepath)
	*outputFilepath = strings.TrimSpace(*outputFilepath)

	var (
		reader io.Reader
		writer io.Writer
		err    error
	)

	if *inputExpression != "" && *inputFilepath != "" {
		return errors.New("you have specified too many inputs")
	} else if *inputExpression != "" && *inputFilepath == "" {
		reader = strings.NewReader(*inputExpression)
	} else if *inputExpression == "" && *inputFilepath != "" {
		reader, err = os.Open(*inputFilepath)
		if err != nil {
			return err
		}
		defer closeFile(reader.(*os.File))
	} else {
		return errors.New("you have not specified any input")
	}

	if *outputFilepath != "" {
		writer, err = os.Create(*outputFilepath)
		if err != nil {
			return err
		}
		defer closeFile(writer.(*os.File))
	} else {
		writer = os.Stdout
	}

	chPtr := &lab2.ComputeHandler{
		Reader: reader,
		Writer: writer,
	}

	if errCompute := chPtr.Compute(); errCompute != nil {
		return errCompute
	}

	return nil
}

func main() {
	if err := processExpression(); err != nil {
		log.Fatal(err)
	}
}
