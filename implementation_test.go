package lab2

import (
	"fmt"
	"os"
	"testing"

	"gopkg.in/check.v1"
)

type verifyCase struct {
	input    string
	expected string
}

func runTestCases(cases *[]verifyCase, c *check.C) {
	for _, testCase := range *cases {
		res, err := PrefixToInfix(testCase.input)
		if err != nil {
			c.Fail()
		} else {
			c.Assert(res, check.Equals, testCase.expected)
		}
	}
}

type TestSuite struct{}

func TestImplementation(t *testing.T) {
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

	check.Run(&TestSuite{}, conf)
}

func (s *TestSuite) TestSimpleExpressions(c *check.C) {
	cases := []verifyCase{
		{
			input:    "^ 2 3",
			expected: "2 ^ 3",
		},
		{
			input:    "/ 8 4",
			expected: "8 / 4",
		},
		{
			input:    "+ 2 2",
			expected: "2 + 2",
		},
		{
			input:    "- 100 A",
			expected: "100 - A",
		},
		{
			input:    "* 2 ^ 2 c",
			expected: "2 * 2 ^ c",
		},
		{
			input:    "^ * 2 2 2",
			expected: "(2 * 2) ^ 2",
		},
		{
			input:    "- + 3 3 20",
			expected: "3 + 3 - 20",
		},
		{
			input:    "* / 6 B + 1 2",
			expected: "6 / B * (1 + 2)",
		},
		{
			input:    "+ * 4 ^ 8 2 11",
			expected: "4 * 8 ^ 2 + 11",
		},
		{
			input:    "+ 47 / * a b 2",
			expected: "47 + a * b / 2",
		},
		{
			input:    "+ 5 * - 4 L 3",
			expected: "5 + (4 - L) * 3",
		},
		{
			input:    "/ / / 12 12 12 12",
			expected: "12 / 12 / 12 / 12",
		},
		{
			input:    "+ / 99 + 51 1 12",
			expected: "99 / (51 + 1) + 12",
		},
	}

	runTestCases(&cases, c)
}

func (s *TestSuite) TestBigExpressions(c *check.C) {
	cases := []verifyCase{
		{
			input:    "+ - 10 * 7 + 3 2 ^ 7 2",
			expected: "10 - 7 * (3 + 2) + 7 ^ 2",
		},
		{
			input:    "/ * / 22 12 44 * 1 + 10 1",
			expected: "22 / 12 * 44 / (1 * (10 + 1))",
		},
		{
			input:    "- - + 12 111 + 12 1 - 10 2",
			expected: "12 + 111 - (12 + 1) - (10 - 2)",
		},
		{
			input:    "^ 55 + - + * 12 12 / 12 12 1 2",
			expected: "55 ^ (12 * 12 + 12 / 12 - 1 + 2)",
		},
		{
			input:    "+ + ^ 12 1 / 122 1 * * 11 2 - 1 0",
			expected: "12 ^ 1 + 122 / 1 + 11 * 2 * (1 - 0)",
		},
		{
			input:    "- + a / b 2 / - + 1 b / sin n 2 3",
			expected: "a + b / 2 - (1 + b - sin(n) / 2) / 3",
		},
		{
			input:    "- + 5 / 9 2 / - + 1 7 / + 8 4 2 3",
			expected: "5 + 9 / 2 - (1 + 7 - (8 + 4) / 2) / 3",
		},
	}

	runTestCases(&cases, c)
}

func (s *TestSuite) TestLargeExpressions(c *check.C) {
	cases := []verifyCase{
		{
			input:    "- + - ^ 2 + 4 12 3 500 - + G ^ / 12 100 2 B",
			expected: "2 ^ (4 + 12) - 3 + 500 - (G + (12 / 100) ^ 2 - B)",
		},
		{
			input:    "- L - ^ 5 - 8 7 - + O / 9 2 + 2 A",
			expected: "L - (5 ^ (8 - 7) - (O + 9 / 2 - (2 + A)))",
		},
		{
			input:    "- - A + ^ - 10 + 3 b - 1 L 20 ^ 11 / 22 2",
			expected: "A - ((10 - (3 + b)) ^ (1 - L) + 20) - 11 ^ (22 / 2)",
		},
		{
			input:    "- + + - - + F C * / 100 2 20 A ^ B C 2 ^ + 0 P Z",
			expected: "F + C - 100 / 2 * 20 - A + B ^ C + 2 - (0 + P) ^ Z",
		},
		{
			input:    "^ - + / 10 + - 20 A H B - + L 11 d + 20 * / L 200 M",
			expected: "(10 / (20 - A + H) + B - (L + 11 - d)) ^ (20 + L / 200 * M)",
		},
	}

	runTestCases(&cases, c)
}

func (s *TestSuite) TestInvalidInput(c *check.C) {
	cases := map[string]error{
		"":                              errParse,
		"         ":                     errParse,
		"1 345 66 3 1 5":                errParse,
		"/ / / * / 22 12 44 * 1 + 10 1": errParse,
		"/ / / 12 12 12 12 12 12 12":    errParse,
		"1 ? + 2 3":                     errParse,
		"-------- *** +++ / ^":          errParse,
		"+ + + + + + +":                 errParse,
		"& 1 2":                         errUnknownOperator,
		"? ! sdfsdf --  ?? %$! * * )":   errUnknownOperator,
		"& / 99 + 51 1 12":              errUnknownOperator,
		"= / 99 + 51 1 12":              errUnknownOperator,
	}

	for input, expected := range cases {
		_, err := PrefixToInfix(input)
		c.Assert(err, check.DeepEquals, expected)
	}
}

func ExamplePrefixToInfix() {
	res, err := PrefixToInfix("- + - 6 - 4 ^ A 2 12 - 2 1")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Println(res)
	}

	// Output:
	// 6 - (4 - A ^ 2) + 12 - (2 - 1)
}
