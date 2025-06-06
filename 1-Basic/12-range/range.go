package main

import "fmt"

func main() {
	message := "Hello World"

	// range works on copy of data structure
	// so changes value inside loop using range does not chat the orignal value

	for i, v := range message {
		// fmt.Println(i, v)
		fmt.Printf("Index: %d, Rune: %c\n", i, v)
	}
}
