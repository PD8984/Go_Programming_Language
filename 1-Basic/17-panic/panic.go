package main

import "fmt"

func main() {
	// panic(interface{})

	// example of valid input
	process(10)

	// example of invalid input
	process(-3)
}

func process(input int) {
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	if input < 0 {
		fmt.Println("Before Panic")
		panic("Input must be a non-negative number")
		// defer fmt.Println("After Panic") // Anything after panic in unreachable after runtime
	}
	fmt.Println("Processing input:", input)
}
