package maps

import (
	"fmt"
)

func LearnMaps() {
	fmt.Println("---------- maps ----------")

	// map is also a pointer, and is reference pass type

	// declare a string-int map
	var myMap1 map[string]int

	if myMap1 == nil {
		fmt.Println("myMap1 is null")
	}

	myMap1 = make(map[string]int, 10)
	myMap1["Man"] = 0
	myMap1["Woman"] = 1
	myMap1["Jet"] = 2

	fmt.Println(myMap1)

	// declare without giving size, it will auto set size
	myMap2 := make(map[string]string)
	myMap2["Man"] = "0"
	myMap2["Woman"] = "1"
	myMap2["Jet"] = "2"
	fmt.Println(myMap2)

	// declare and set value
	myMap3 := map[string]string{
		"one": "java",
		"two": "go",
		// note that even the last line still needs comma
		"three": "cpp",
	}
	fmt.Println(myMap3)

	cityMap := make(map[string]string)

	// add entry to map
	cityMap["China"] = "BeiJing"
	cityMap["Japan"] = "Tokyo"
	cityMap["America"] = "NewYorkCity"

	// traverse map
	for k, v := range cityMap {
		fmt.Println("k = ", k, ", v = ", v)
	}

	// delete entry
	delete(cityMap, "Japan")

	// modify entry
	cityMap["America"] = "DC"

	for k, v := range cityMap {
		fmt.Println("k = ", k, ", v = ", v)
	}

	// see whether the map contains key
	element, contains := cityMap["China"]
	if contains {
		fmt.Println("Element is in the map, value: ", element)
	} else {
		fmt.Println("Element is not in the map")
	}

	fmt.Println("")
}