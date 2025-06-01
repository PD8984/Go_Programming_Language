package main

import "fmt"

func main() {
	// fmt.Println("Hello" + "World")
	// fmt.Println("9 X 10 =", 9*10)
	// fmt.Println("100.18/2 = ", 100.18/2)
	// fmt.Println(true && false)
	// fmt.Println(true || false)
	// fmt.Println(!true)

	var age int
	var name string = "Preet"
	var name1 = "Aman"
	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(name1)

	count := 10
	lastname := "Dhiman"
	fmt.Println(count)
	fmt.Println(lastname)

	// Default Values
	// Numeric Types: 0
	// Boolean Types: false
	// String Type: ""
	// Pointers, slices, maps, functions, and structs: nil

	// ---- SCOPE
	printname("Preet", "Dhiman")

}

func printname(firstName string, lastname string) {
	fmt.Println(firstName, lastname)
}
