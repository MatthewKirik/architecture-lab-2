package lab2

// import (
// 	"bytes"
// 	"strings"
// 	"testing"

// 	. "gopkg.in/check.v1"
// )

// func TestCompute(t *testing.T) { TestingT(t) }

// func (s *TestSuite) TestComputeOutput(c *C) {
// 	inputStr, expected := "+55-2", "5 + 5 - 2"
// 	buf := new(bytes.Buffer)
// 	reader := strings.NewReader(inputStr)
// 	handler := ComputeHandler{Reader: reader, Writer: buf}
// 	handler.Compute()
// 	c.Assert(buf.String(), Equals, expected)
// }

// func (s *TestSuite) TestComputeSyntax(c *C) {
// 	errorExamples := map[string]ComputeHandler{
// 		"input is not specified":   ComputeHandler{},
// 		"output is not specified":  ComputeHandler{Reader: strings.NewReader("+ 2 4")},
// 		"invalid input expression": ComputeHandler{Reader: strings.NewReader(""), Writer: new(bytes.Buffer)},
// 	}
// 	for expected, obtained := range errorExamples {
// 		errObtained := obtained.Compute()
// 		c.Assert(errObtained, ErrorMatches, expected)
// 	}
// }
