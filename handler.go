package lab2

import (
	"io"
	"io/ioutil"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Reader io.Reader
	Writer io.Writer
}

func (ch *ComputeHandler) Compute() error {
	// strings.NewReader(...) -> [-e flag]
	// os.Open(...) -> [-f flag]
	// os.Create(...) -> [-o flag]

	prefixStr, errRead := ioutil.ReadAll(ch.Reader)
	if errRead != nil {
		return errRead
	}

	resExpr, errCalc := PrefixToInfix(string(prefixStr))
	if errCalc != nil {
		return errCalc
	}

	textBytes := []byte(resExpr)
	_, errWrite := ch.Writer.Write(textBytes)
	if errWrite != nil {
		return errWrite
	}

	return nil
}
