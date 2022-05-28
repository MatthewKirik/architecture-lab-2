package lab2

import (
	"bytes"
	// "strings"
	"testing"
	// "io"

	. "gopkg.in/check.v1"
)

type mockReader struct {
	ReadWasCalled bool
	DummyInput    []byte
}

func (mr *mockReader) Read(p []byte) (int, error) {
	mr.ReadWasCalled = true
	copy(p, mr.DummyInput)

	var bytesWasRead int
	if len(p) < len(mr.DummyInput) {
		bytesWasRead = len(p)
	} else {
		bytesWasRead = len(mr.DummyInput)
	}

	// can't we go the way below instead
	// of such long confitional sentence?
	// bytesWasRead := len(string(p))

	return bytesWasRead, nil
}

type mockWriter struct {
	WriteWasCalled bool
	DummyOutput    *bytes.Buffer
}

func (mw *mockWriter) Write(p []byte) (int, error) {
	mw.WriteWasCalled = true

	nb, err := mw.DummyOutput.Write(p)
	if err != nil {
		return 0, err
	}

	return nb, nil
}

func TestCompute(t *testing.T) {
	TestingT(t)
}

func (s *TestHandlerSuite) TestReadWasCalled(c *C) {
	// TODO: implements
}

func (s *TestHandlerSuite) TestWriteWasCalled(c *C) {
	// TODO: implements
}

func (s *TestHandlerSuite) TestInputSyntaxError(c *C) {
	// TODO: implements
}

func (s *TestHandlerSuite) TestInputMatchesOutput(c *C) {
	// TODO: implements
}

// func (s *TestSuite) TestComputeOutput(c *C) {
//     inputStr, expected := "+55-2", "5 + 5 - 2"
//     buf := new(bytes.Buffer)
//     reader := strings.NewReader(inputStr)
//     handler := ComputeHandler{Reader: reader, Writer: buf}
//     handler.Compute()
//     c.Assert(buf.String(), Equals, expected)
// }

// func (s *TestSuite) TestComputeSyntax(c *C) {
//     errorExamples := map[string]ComputeHandler{
//         "input is not specified":   ComputeHandler{},
//         "output is not specified":  ComputeHandler{Reader: strings.NewReader("+ 2 4")},
//         "invalid input expression": ComputeHandler{Reader: strings.NewReader(""), Writer: new(bytes.Buffer)},
//     }
//     for expected, obtained := range errorExamples {
//         errObtained := obtained.Compute()
//         c.Assert(errObtained, ErrorMatches, expected)
//     }
// }
