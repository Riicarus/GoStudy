package typecast

import (
	"strconv"
	"fmt"
)

func LearnTypeCast() {
	// type cast in go must be explicitly declared

	fmt.Println("---------- type cast ----------")

	var aInt int = 100
	var aFloat float32 = 3.3
	res := aFloat / float32(aInt)
	fmt.Println(res)

	// string to int
	aStr := "100"
	bInt, err := strconv.Atoi(aStr)

	if err == nil {
		fmt.Printf("aStr: %T %s, bInt: %T %d\n", aStr, aStr, bInt, bInt)
	} else {
		fmt.Printf("err: %s\n", err)
	}

	// int to string
	cInt := 200
	bStr := strconv.Itoa(cInt)
	fmt.Printf("bStr: %T %s, cInt: %T %d\n", bStr, bStr, cInt, cInt)

	fmt.Println("")
}