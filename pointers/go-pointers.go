package pointers

import (
	"fmt"
)

func LearnPointers() {
	fmt.Println("---------- pointers ----------")
	oldVal := 1
	fmt.Println("oldVal = ", oldVal)
	changeValue(&oldVal)
	fmt.Println("newVal = ", oldVal)
	fmt.Println("")
}

func changeValue(p *int) {
	*p = 10
}