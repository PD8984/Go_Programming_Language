package main

import "fmt"

func main() {
	// func name(parameters list){
	// 	//code block
	//  //return
	// }

	// declaration of function inside main another function is not allowed in go

	fmt.Println(addition(2, 3))

	greet := func() {
		fmt.Println("Hello Anonymous function")
	}
	greet()

	operation := addition
	result := operation(2, 5)
	fmt.Println(result)
}
func addition(num1, num2 int) int {
	return num1 + num2
}
