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

func (mw *mockReader) Read(p []byte) (int, error) {
	mw.ReadWasCalled = true
	copy(p, mw.DummyInput)

	var bytesWasRead int
	if len(p) < len(mw.DummyInput) {
		bytesWasRead = len(p)
	} else {
		bytesWasRead = len(mw.DummyInput)
	}

	// can't we go the way below intead
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
