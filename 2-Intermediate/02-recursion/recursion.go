package main

import "fmt"

func main() {
	fmt.Println(factorial(10))
	fmt.Println(factorial(20))

	fmt.Println(sumOfDigits(111))
	fmt.Println(sumOfDigits(123))
}

func factorial(n int) int {
	//Base case: factorial of 0 is 1
	if n == 0 {
		return 1
	}
	// Recursive case: factorial of n is n * factorial(n-1)
	return n * factorial(n-1)
}

func sumOfDigits(n int) int {
	// base case
	if n < 10 {
		return n
	}
	return n%10 + sumOfDigits(n/10)
}
