package lab2

import (
	"bytes"
	"os"
	"testing"

	"gopkg.in/check.v1"
)

type mockReader struct {
	ReadWasCalled bool
	DummyInput    *bytes.Buffer
}

func (mr *mockReader) Read(p []byte) (int, error) {
	mr.ReadWasCalled = true
	bytesWasRead, err := mr.DummyInput.Read(p)

	return bytesWasRead, err
}

type mockWriter struct {
	WriteWasCalled bool
	DummyOutput    *bytes.Buffer
}

func (mw *mockWriter) Write(p []byte) (int, error) {
	mw.WriteWasCalled = true
	bytesWasWritten, err := mw.DummyOutput.Write(p)

	return bytesWasWritten, err
}

type TestHandlerSuite struct{}

func TestCompute(t *testing.T) {
	conf := &check.RunConf{
		Output:        os.Stdout,
		Stream:        false,
		Verbose:       false,
		Filter:        "",
		Benchmark:     false,
		BenchmarkTime: 0,
		BenchmarkMem:  false,
		KeepWorkDir:   false,
	}

	check.Run(&TestHandlerSuite{}, conf)
}

func (s *TestHandlerSuite) TestReadWriteWasCalled(c *check.C) {
	// consider use s.dummyObj interface{} ???
	dummyStr := "Hello, World!"
	inputBuf := []byte(dummyStr)
	outputBuf := make([]byte, 0, 64)
	mr := &mockReader{
		ReadWasCalled: false,
		DummyInput:    bytes.NewBuffer(inputBuf),
	}
	mw := &mockWriter{
		WriteWasCalled: false,
		DummyOutput:    bytes.NewBuffer(outputBuf),
	}

	handler := ComputeHandler{
		Reader: mr,
		Writer: mw,
	}

	inputLenBeforeTest := mr.DummyInput.Len()
	outputLenBeforeTest := mw.DummyOutput.Len()

	handler.Compute()

	inputLenAfterTest := mr.DummyInput.Len()
	outputLenAfterTest := mw.DummyOutput.Len()

	wasBytesRead := inputLenBeforeTest > inputLenAfterTest
	wasBytesWritten := outputLenBeforeTest < outputLenAfterTest

	c.Assert(mr.ReadWasCalled, check.Equals, true)
	c.Assert(wasBytesRead, check.Equals, true)
	c.Assert(mw.WriteWasCalled, check.Equals, true)
	c.Assert(wasBytesWritten, check.Equals, true)
}

// func (s *TestHandlerSuite) TestWriteWasCalled(c *C) {
// 	dummyStr := "Hello, World!"
// 	mr := &mockReader{
// 		ReadWasCalled: false,
// 		DummyInput:    []byte(dummyStr),
// 	}
// 	mw := &mockWriter{
// 		WriteWasCalled: false,
// 		DummyOutput:    []byte(dummyStr),
// 	}

func (s *TestHandlerSuite) TestWriteWasCalled(c *check.C) {
	// TODO: implements
}

func (s *TestHandlerSuite) TestInputSyntaxError(c *check.C) {
	// TODO: implements
}

func (s *TestHandlerSuite) TestInputMatchesOutput(c *check.C) {
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
