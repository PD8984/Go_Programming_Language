package main

import "fmt"

func main() {
	// var arrayName [size]elementType;
	var number [5]int
	fmt.Println(number)

	number[0] = 10
	number[1] = 23
	number[2] = 34
	number[3] = 45
	number[4] = 56
	fmt.Println(number)

	fruits := [4]string{"Apple", "Orange", "Banana", "Grapes"}
	fmt.Println(fruits)

	for index, value := range fruits {
		fmt.Println(index, value)
	}

	fmt.Println("\nFor...range loop (value only):")
	for _, value := range fruits { // Use '_' to ignore the index
		fmt.Println("Value:", value)
	}

	fmt.Println("\nFor...range loop (index only):")
	for index := range fruits { // Only specify 'index'
		fmt.Println("Index:", index)
	}

	fmt.Println("Lenths of array in fruits are ", len(fruits))

	// coping array
	fruitBasket := fruits
	fmt.Println(fruitBasket)

}
