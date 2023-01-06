package slices

import (
	"fmt"
)

func LearnSlices() {
	fmt.Println("---------- arrays and slices ----------")
	// slice in go is a dynamic array
	// slice doesn't declare length
	// slice is a referance pass in func, pass a pointer to func
	// slice is a pointer points to the first element of one dynamic array
	// declare slice
	dynamicArray := []int{1, 2, 3, 4}

	// can use make to delay slice's init
	slice1 := make([]int, 3)
	slice1[0] = 100
	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

	var slice2 []int

	// judge whether a slice is empty(has no element or not inited)
	if slice2 == nil {
		fmt.Println("slice2 is empty")
	} else {
		fmt.Println("slice2 has element")
	}

	numbers := make([]int, 3, 5)

	// if want to add one element to index out of len, we should use append() to add element to index = len
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
	numbers = append(numbers, 6)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
	numbers = append(numbers, 8)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// if want to add one element to index out of cap, it will auto resize slice cap to len << 2
	numbers = append(numbers, 10)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// fixed length array
	// is a value pass in func
	// array of different length is different type, can not change between each other
	array := [10]int{1, 2, 3, 4}

	for i := 0; i < len(array); i++ {
		fmt.Println("fixed length array[", i, "] = ", array[i])
	}

	fmt.Println("use range")
	for index, value := range array {
		fmt.Println("index = ", index, ", value = ", value)
	}

	// yet another use of range 
	for idx, char := range "hello" {
		fmt.Println("idx = ", idx, ", char = ", char)
	}
	
	fmt.Println("use slice")
	for _, value := range dynamicArray {
		fmt.Println("slice values = ", value)
	}

	// sub slice, use the same location with sup slice, like List.toArray()
	// [start, end)
	// [: end] ==> [0, end)
	// [start :] ==> [start, last)
	subNumbers := numbers[1:4]
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(subNumbers), cap(subNumbers), subNumbers)

	// use copy to create a new location for sub slice
	subNumbers2 := make([]int, 4)
	copy(subNumbers2, subNumbers)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(subNumbers2), cap(subNumbers2), subNumbers2)


	fmt.Println("")
}