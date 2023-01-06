package functions

import (
	"fmt"
)

func LearnFunctions() {
	fmt.Println("---------- functions ----------")
	fmt.Println("func foo ret = ", foo("hello", 10))
	ret1, ret2 := fooMulti("hello", 10)
	fmt.Println("func fooMulti ret1 = ", ret1, " ret2 = ", ret2)
	ret3, ret4, ret5 := fooMultiWithName("hello", 10, 100)
	fmt.Println("func fooMultiWithName ret3 = ", ret3, " ret4 = ", ret4, " ret5 = ", ret5)
	fmt.Println("")
}

func foo(str string, integer int) int {
	fmt.Println("str = ", str)
	fmt.Println("integer = ", integer)

	ret := 100

	return ret
}

func fooMulti(str string, integer int) (string, int) {
	return str, integer
}

func fooMultiWithName(str string, i1 int, i2 int) (strRet string, r1 int, r2 int) {
	return str, i1, i2
}