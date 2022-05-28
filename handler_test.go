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

func createMocks(inputStr string) (*mockReader, *mockWriter) {
	inputBuf := []byte(inputStr)
	outputBuf := make([]byte, 0, 64)
	mr := &mockReader{
		ReadWasCalled: false,
		DummyInput:    bytes.NewBuffer(inputBuf),
	}
	mw := &mockWriter{
		WriteWasCalled: false,
		DummyOutput:    bytes.NewBuffer(outputBuf),
	}

	return mr, mw
}

type TestCase struct {
	InputStr        string
	ExpectedStr     string
	IsErrorExpected bool
}

type TestHandlerSuite struct {
	mr *mockReader
	mw *mockWriter
	ch *ComputeHandler
	// testCases *[]TestCase
}

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

	dummyStr := "Hello, World!"
	mr, mw := createMocks(dummyStr)
	ch := &ComputeHandler{
		Reader: mr,
		Writer: mw,
	}

	// testCases := &[]TestCase{
	// 	TestCase{
	// 		InputStr: ,
	// 		ExpectedStr: ,
	// 		IsErrorExpected: ,
	// 	}
	// }

	suite := &TestHandlerSuite{
		mr,
		mw,
		ch,
	}
	check.Run(suite, conf)
}

func (s *TestHandlerSuite) SetUpTest(c *check.C) {
	// consider use s.dummyObj interface{} ???
	dummyStr := "Dummy input string"
	mr, mw := createMocks(dummyStr)
	ch := &ComputeHandler{
		Reader: mr,
		Writer: mw,
	}

	s.mr = mr
	s.mw = mw
	s.ch = ch
}

func (s *TestHandlerSuite) TearDownTest(c *check.C) {
	s.mr = nil
	s.mw = nil
	s.ch = nil
}

func (s *TestHandlerSuite) TestReadWasCalled(c *check.C) {
	inputLenBeforeTest := s.mr.DummyInput.Len()

	err := s.ch.Compute()

	inputLenAfterTest := s.mr.DummyInput.Len()
	wasBytesRead := inputLenBeforeTest > inputLenAfterTest

	if err != nil {
		c.Assert(err, check.NotNil)
	} else {
		c.Assert(s.mr.ReadWasCalled, check.Equals, true)
		c.Assert(wasBytesRead, check.Equals, true)
	}
}

func (s *TestHandlerSuite) TestWriteWasCalled(c *check.C) {
	outputLenBeforeTest := s.mw.DummyOutput.Len()

	err := s.ch.Compute()

	outputLenAfterTest := s.mw.DummyOutput.Len()
	wasBytesWritten := outputLenBeforeTest < outputLenAfterTest

	if err != nil {
		c.Assert(err, check.NotNil)
	} else {
		c.Assert(s.mw.WriteWasCalled, check.Equals, true)
		c.Assert(wasBytesWritten, check.Equals, true)
	}
}

func (s *TestHandlerSuite) TestInput(c *check.C) {
	testCases := []TestCase{
		{
			InputStr:        "ab rakada bra12 123 ? + /",
			ExpectedStr:     "",
			IsErrorExpected: true,
		},
		{
			InputStr:        "- + 5 / 9 2 / - + 1 7 / + 8 4 2 3",
			ExpectedStr:     "5 + 9 / 2 - (1 + 7 - (8 + 4) / 2) / 3",
			IsErrorExpected: false,
		},
		{
			InputStr:        "- + 1 / 5 6 + 1 2",
			ExpectedStr:     "1 + 5 / 6 - (1 + 2)",
			IsErrorExpected: false,
		},
		{
			InputStr:        "- - + * / / + - *",
			ExpectedStr:     "",
			IsErrorExpected: true,
		},
		{
			InputStr:        "- ? ----- --? a! # $ ; !!!!! ++ /",
			ExpectedStr:     "",
			IsErrorExpected: true,
		},
	}

	for _, testCase := range testCases {
		mr, mw := createMocks(testCase.InputStr)
		s.ch = &ComputeHandler{
			Reader: mr,
			Writer: mw,
		}

		err := s.ch.Compute()

		if testCase.IsErrorExpected {
			c.Assert(err, check.NotNil)
		} else {
			actualWritten := mw.DummyOutput.String()
			c.Assert(testCase.ExpectedStr, check.Equals, actualWritten)
		}
	}
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
