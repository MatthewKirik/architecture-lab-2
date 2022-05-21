package lab2

import (
	"fmt"

	. "gopkg.in/check.v1"
)

func (s *TestSuite) TestPrefixToInfix(c *C) {
	samples := map[string]string{
		"/ * / 22 12 44 * 1 + 10 1":               "22 / 12 * 44 / (1 * (10 + 1))",
		"/ / / 12 12 12 12":                       "12 / 12 / 12 / 12",
		"+ / 99 + 51 1 12":                        "99 / (51 + 1) + 12",
		"+ + ^ 12 1 / 122 1 * * 11 2 - 1 0":       "12^1 + 122 / 1 + 11 * 2 * (1 - 0)",
		"^ 55 + - + * 12 12 / 12 12 1 2":          "55 ^ (12 * 12 + 12 / 12 - 1 + 2)",
		"- - + 12 111 + 12 1 - 10 2":              "((12+111)) - (12 + 1) - (10 - 2)",
		"- + 5 / 9 2 / - + 1 7 / + 8 4 2 3":       "5 + (9 / 2) - ((1 + 7) - (8 + 4) / 2) / 3",
		"- + a / b 2 / - + 1 b / + someVar n 2 3": "a + b / 2 - (1 + b - (someVar + n) / 2) / 3",
		"- + a / b 2 / - + 1 b / sin n 2 3":       "a + b / 2 - (1 + b - sin(n) / 2) / 3",
	}
	for prefix, expected := range samples {
		res, err := PrefixToInfix(prefix)
		if err != nil {
			c.Assert(err, ErrorMatches, expected)
		} else {
			c.Assert(res, Equals, expected)
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
