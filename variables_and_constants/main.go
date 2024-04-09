package main

import (
	"fmt"
	"reflect"
)

// Variables are spaces in computer's memory to store values 
// that can be modified during the execution of programs.

func main() {
	
	// explicity types declared variable
	var explicity string = "Shivam Singh"
	fmt.Printf("explicity type variable: %v\n", explicity)
	// find which type
	fmt.Printf("explicity type variable: %v\n", reflect.TypeOf(explicity))

	// inferred types declared variable
	inferred := "Shivam Singh"
	fmt.Printf("inferred type variable: %v\n", inferred)
	// find which type
	fmt.Printf("inferred type variable: %v\n", reflect.TypeOf(inferred))
}
