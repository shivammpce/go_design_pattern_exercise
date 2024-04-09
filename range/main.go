package main

import "fmt"

// We use the 'range' keyword to retrieve every item in a collection
func main() {
	my_array := []int{1, 2, 3, 4, 5}

	// using range
	for index, value := range my_array {
		fmt.Printf("First Index is %d and Value is %d\n", index, value)
	}

	// using for
	// The 'len' method is used to know the length of collection
	for index := 0; index < len(my_array); index++ {
		value := my_array[index]
		fmt.Printf("Second Index is %d and Value is %d\n", index, value)
	}

}
