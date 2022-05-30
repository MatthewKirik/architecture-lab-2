package lab2

import (
	"io"
	"io/ioutil"
)

type ComputeHandler struct {
	Reader io.Reader
	Writer io.Writer
}

func (ch *ComputeHandler) Compute() error {
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
