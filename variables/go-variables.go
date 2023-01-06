package variables

import (
	"fmt"
)

// global variable
var e float32 = 1.1
var f = false

// const variable
const length = 10

// iota: increase 1 each line
// iota can only be used with const to imply the enum type
const (
	BEJING = iota * 10
	SHANGHAI
	SHENZHEN
	CHENGDU = 3
)

const (
	o, p = iota + 1, iota + 2 // iota = 0
	q, r                      // iota = 1

	s, t = iota * 2, iota * 3 // iota = 2
)

func LearnVariables() {
	fmt.Println("---------- variables ----------")

	// local variable
	var a int = 10
	b := 20
	c, d := "string", true

	fmt.Println("Local Variable")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)
	fmt.Println("")

	fmt.Println("Global Variable")
	fmt.Println("e = ", e)
	fmt.Println("f = ", f)
	fmt.Println("")

	fmt.Println("Const Variable")
	fmt.Println("length = ", length)
	fmt.Println("")

	fmt.Println("Const Variable With IOTA")
	fmt.Println("BEIJING = ", BEJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)
	fmt.Println("CHENGDU = ", CHENGDU)
	fmt.Println("o = ", o)
	fmt.Println("p = ", p)
	fmt.Println("q = ", q)
	fmt.Println("r = ", r)
	fmt.Println("s = ", s)
	fmt.Println("t = ", t)
	fmt.Println("")
}