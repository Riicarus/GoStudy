package interfaces

import (
	"fmt"
)

/*
type interface_name interface {
	method_name_1 [return_type]
	method_name_2 [return_type]
	// ...
}

type struct_name struct {
	// variables
}

func (struct_name_variable struct_name) method_name_1 [return_type] {
	// ...
}

func (struct_name_variable struct_name) method_name_2 [return_type] {
	// ...
}
*/
type CodeHelper interface {
	Help() interface{}
}

type CodeRunner interface {
	RunCode() interface{}
}

type JavaCodeHelper struct {
	name string
}

func (helper *JavaCodeHelper) Help() interface{} {
	fmt.Printf("%s help java code\n", helper.name)

	return helper
}

func (helper *JavaCodeHelper) RunCode() interface{} {
	fmt.Printf("%s run java code\n", helper.name)

	return helper
}

type GoCodeHelper struct {
	name string
}

func (helper *GoCodeHelper) Help() interface{} {
	fmt.Printf("%s help go code\n", helper.name)

	return helper
}

func (helper *GoCodeHelper) RunCode() interface{} {
	fmt.Printf("%s run go code\n", helper.name)

	return helper
}

func LearnInterfaces() {
	fmt.Println("---------- interfaces ----------")

	var codeHelper CodeHelper

	codeHelper = &JavaCodeHelper{"JavaCodeHelper"}
	codeHelper.Help()

	// pair<type: "GoCodeHelper", value: &GoCodeHelper{}>
	codeHelper = &GoCodeHelper{"GoCodeHelper"}
	codeHelper.Help()

	// universal interface
	// we can use it to state any type of variable, usage: var a interface{} = ...
	var allType interface{} = codeHelper
	
	// type assert for universal interface
	ch, ok := allType.(CodeHelper)
	if ok {
		fmt.Println("this is a CodeHelper")
		ch.Help()
	} else {
		fmt.Println("this is not a CodeHelper")
	}

	// phlymorphic type cast: interface_name_2_variable, couldCast := interface_name_1_variable.(interface_name_2)
	// polymorphism(多态) can be used between interfaces, but not no-interfaces
	// there's a pair of type and value passing between variables
	// pair<type, value>, type could be static type like int, string .etc or concrete type like struct, interface
	// pair of the same vairable during passing would never change, it's the basic of type assert and polymorphism

	// pair<type: , value: >
	var codeRunner CodeRunner

	// pair<type: "GoCodeHelper", value: &GoCodeHelper{}>
	codeRunner, _ = codeHelper.(CodeRunner)
	codeRunner.RunCode()

	goCodeHelper := GoCodeHelper{"GoHelper"}
	goCodeHelper.Help()
	goCodeHelper.RunCode()


	fmt.Println("")
}