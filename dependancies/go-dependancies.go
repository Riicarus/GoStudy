package dependancies

import (
	"fmt"
	_ "github.com/riicarus/gostudy/dependancies/lib1"
	mylib2 "github.com/riicarus/gostudy/dependancies/lib2"
)

func LearnDependancies() {
	fmt.Println("---------- dependancies ----------")
	// lib1.Lib1Test()
	mylib2.Lib2Test()
	fmt.Println("")
}