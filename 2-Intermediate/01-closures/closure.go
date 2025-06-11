package main

import "fmt"

func main() {
	// sequence := adder()
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	//
	// sequence2 := adder()
	// fmt.Println(sequence2())

	subtractor := func() func(int) int {
		countdown := 99
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()
	// using the closure subtractor
	fmt.Println(subtractor(1))
	fmt.Println(subtractor(2))
	fmt.Println(subtractor(3))
	fmt.Println(subtractor(4))
	fmt.Println(subtractor(1))
}

func adder() func() int {
	i := 0
	fmt.Println("previous value of i:", i)
	return func() int {
		i++
		fmt.Println("added 1 to i")
		return i
	}
}
