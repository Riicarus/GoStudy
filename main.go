package main

import (
	"fmt"
	variables "github.com/riicarus/gostudy/variables"
	functions "github.com/riicarus/gostudy/functions"
	dependancies "github.com/riicarus/gostudy/dependancies"
	pointers "github.com/riicarus/gostudy/pointers"
	slices "github.com/riicarus/gostudy/slices"
	maps "github.com/riicarus/gostudy/maps"
	structs "github.com/riicarus/gostudy/structs"
	typecast "github.com/riicarus/gostudy/typecast"
	interfaces "github.com/riicarus/gostudy/interfaces"
	reflects "github.com/riicarus/gostudy/reflects"
	routines "github.com/riicarus/gostudy/routines"
)

func main() {
	fmt.Println("Study Go From Here")

	variables.LearnVariables()

	functions.LearnFunctions()

	dependancies.LearnDependancies()

	pointers.LearnPointers()

	slices.LearnSlices()

	maps.LearnMaps()
	
	structs.LearnStructs()
	
	typecast.LearnTypeCast()
	
	interfaces.LearnInterfaces()

	reflects.LearnReflect()

	routines.LearnRoutines()
}