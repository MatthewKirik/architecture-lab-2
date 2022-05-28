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

func (s *TestSuite) TestPrefixToInfix(c *check.C) {
	samples := map[string]string{
		"/ * / 22 12 44 * 1 + 10 1":         "22 / 12 * 44 / (1 * (10 + 1))",
		"/ / / 12 12 12 12":                 "12 / 12 / 12 / 12",
		"+ / 99 + 51 1 12":                  "99 / (51 + 1) + 12",
		"+ + ^ 12 1 / 122 1 * * 11 2 - 1 0": "12 ^ 1 + 122 / 1 + 11 * 2 * (1 - 0)",
		"^ 55 + - + * 12 12 / 12 12 1 2":    "55 ^ (12 * 12 + 12 / 12 - 1 + 2)",
		"- - + 12 111 + 12 1 - 10 2":        "12 + 111 - (12 + 1) - (10 - 2)",
		"- + 5 / 9 2 / - + 1 7 / + 8 4 2 3": "5 + 9 / 2 - (1 + 7 - (8 + 4) / 2) / 3",
		"- + a / b 2 / - + 1 b / sin n 2 3": "a + b / 2 - (1 + b - sin(n) / 2) / 3",
	}
	for prefix, expected := range samples {
		res, err := PrefixToInfix(prefix)
		if err != nil {
			c.Assert(err, check.ErrorMatches, expected)
		} else {
			c.Assert(res, check.Equals, expected)
		}
	}
}

func ExamplePrefixToInfix() {
	res, err := PrefixToInfix("++555")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(res)
	}
}
