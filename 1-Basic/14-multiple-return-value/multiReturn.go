package main

import "fmt"

func main() {
	q, r := divide(3, 6)
	fmt.Println(q)
	fmt.Println(r)
}

// func fucntionName(parameter1 type1, parameter2 type2, ...) (returnType1, returnType2, ...)
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}
