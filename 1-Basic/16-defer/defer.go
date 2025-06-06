package main

import "fmt"

func main() {
	process(10)
}

func process(i int) {
	defer fmt.Println("Deffered i value", i)
	defer fmt.Println("First Deffered statement executed")
	defer fmt.Println("Second Deffered statement executed")
	defer fmt.Println("Third Deffered statement executed")
	i++
	fmt.Println("Normal Execution Statement")
	fmt.Println("value of i", i)
}
