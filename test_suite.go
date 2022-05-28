package lab2

import (
	. "gopkg.in/check.v1"
)

type TestSuite struct{}
type TestHandlerSuite struct{}

func init() {
	Suite(&TestHandlerSuite{})
}
