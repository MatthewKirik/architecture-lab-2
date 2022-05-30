package lab2

import (
	"fmt"
	"os"
	"testing"

	"gopkg.in/check.v1"
)

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

type verifyCase struct {
	input    string
	expected string
}

func (s *TestSuite) TestPrefixToInfixValid(c *check.C) {
	cases := map[string]string{
		"^ 2 3":                             "2 ^ 3",
		"/ 8 4":                             "8 / 4",
		"+ 2 2":                             "2 + 2",
		"- 100 A":                           "100 - A",
		"* 2 ^ 2 c":                         "2 * 2 ^ c",
		"^ * 2 2 2":                         "(2 * 2) ^ 2",
		"- + 3 3 20":                        "3 + 3 - 20",
		"* / 6 B + 1 2":                     "6 / B * (1 + 2)",
		"+ * 4 ^ 8 2 11":                    "4 * 8 ^ 2 + 11",
		"+ 47 / * a b 2":                    "47 + a * b / 2",
		"+ 5 * - 4 L 3":                     "5 + (4 - L) * 3",
		"/ / / 12 12 12 12":                 "12 / 12 / 12 / 12",
		"+ / 99 + 51 1 12":                  "99 / (51 + 1) + 12",
		"+ - 10 * 7 + 3 2 ^ 7 2":            "10 - 7 * (3 + 2) + 7 ^ 2",
		"/ * / 22 12 44 * 1 + 10 1":         "22 / 12 * 44 / (1 * (10 + 1))",
		"- - + 12 111 + 12 1 - 10 2":        "12 + 111 - (12 + 1) - (10 - 2)",
		"^ 55 + - + * 12 12 / 12 12 1 2":    "55 ^ (12 * 12 + 12 / 12 - 1 + 2)",
		"+ + ^ 12 1 / 122 1 * * 11 2 - 1 0": "12 ^ 1 + 122 / 1 + 11 * 2 * (1 - 0)",
		"- + a / b 2 / - + 1 b / sin n 2 3": "a + b / 2 - (1 + b - sin(n) / 2) / 3",
		"- + 5 / 9 2 / - + 1 7 / + 8 4 2 3": "5 + 9 / 2 - (1 + 7 - (8 + 4) / 2) / 3",
		"- + - ^ 2 + 4 12 3 500 - + G ^ / 12 100 2 B":         "2 ^ (4 + 12) - 3 + 500 - (G + (12 / 100) ^ 2 - B)",
		"- L - ^ 5 - 8 7 - + O / 9 2 + 2 A":                   "L - (5 ^ (8 - 7) - (O + 9 / 2 - (2 + A)))",
		"- - A + ^ - 10 + 3 b - 1 L 20 ^ 11 / 22 2":           "A - ((10 - (3 + b)) ^ (1 - L) + 20) - 11 ^ (22 / 2)",
		"- + + - - + F C * / 100 2 20 A ^ B C 2 ^ + 0 P Z":    "F + C - 100 / 2 * 20 - A + B ^ C + 2 - (0 + P) ^ Z",
		"^ - + / 10 + - 20 A H B - + L 11 d + 20 * / L 200 M": "(10 / (20 - A + H) + B - (L + 11 - d)) ^ (20 + L / 200 * M)",
	}

	for input, expected := range cases {
		res, err := PrefixToInfix(input)
		if err != nil {
			c.Fail()
		} else {
			c.Assert(res, check.Equals, expected)
		}
	}
}

func (s *TestSuite) TestPrefixToInfixError(c *check.C) {
	cases := map[string]error{
		"":                              parseErr,
		"         ":                     parseErr,
		"1 345 66 3 1 5":                parseErr,
		"/ / / * / 22 12 44 * 1 + 10 1": parseErr,
		"/ / / 12 12 12 12 12 12 12":    parseErr,
		"1 ? + 2 3":                     parseErr,
		"-------- *** +++ / ^":          parseErr,
		"+ + + + + + +":                 parseErr,
		"& 1 2":                         unknownOperatorErr,
		"? ! sdfsdf --  ?? %$! * * )":   unknownOperatorErr,
		"& / 99 + 51 1 12":              unknownOperatorErr,
		"= / 99 + 51 1 12":              unknownOperatorErr,
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
