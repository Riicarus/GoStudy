package reflects

import (
	"fmt"
	"reflect"
)

func reflectNum(arg interface{}) {
	fmt.Println("type: ", reflect.TypeOf(arg))
	fmt.Println("value: ", reflect.ValueOf(arg))
}

type User struct {
	Id		int
	Name	string
	Age 	int
}

func (user *User) Call() {
	fmt.Printf("user: %v is called...\n", user)
}

func reflectComplexStruct(arg interface{}) {
	fmt.Println(&arg)
	// get type
	inputType := reflect.TypeOf(arg)
	fmt.Println("input type is: ", inputType.Name())

	// get value
	inputValue := reflect.ValueOf(arg)
	fmt.Println("input value is: ", inputValue)

	// get fields
	// 1. get reflect.Type of interface, and get NumField through Type, then traverse NumField
	// 2. get each field, which is data type
	// 3. get value through field's Interface() method
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(1).Interface()

		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// get methods and invoke
	for i := 0; i < inputType.NumMethod(); i++ {
		method := inputType.Method(i)
		fmt.Printf("%s: %v\n", method.Name, method.Type)
	}
}

func reflectMethodOfStructPointer(arg interface{}) {
	argType := reflect.TypeOf(arg)
	for i := 0; i < argType.NumMethod(); i++ {
		method := argType.Method(i)
		fmt.Printf("%s: %v\n", method.Name, method.Type)
	}
}

func LearnReflect() {
	fmt.Println("---------- reflects ----------")

	a := 1.2345
	reflectNum(a)

	user := User{1, "Riicarus", 20}
	var b interface{} = user
	fmt.Println(&user)
	fmt.Println(&b)
	reflectComplexStruct(user)
	reflectMethodOfStructPointer(&user)

	fmt.Println("")
}