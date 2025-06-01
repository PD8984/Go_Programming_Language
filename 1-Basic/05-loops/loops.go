package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/* // for loop */

	// // Simple Interation over a range
	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(i)
	// }
	//
	// // iterate over collection
	// numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// for index, value := range numbers {
	// 	fmt.Printf("Index: %v, Value: %v\n", index, value)
	// }

	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		continue // continue the loop but skip the rest of lines/statement
	// 	}
	// 	fmt.Println("Odd Number", i)
	// 	if i == 5 {
	// 		break // break out of loop
	// 	}
	// }

	// ASTERISK LAYOUT
	// rows := 5
	// // Outer loop
	// for i := 1; i <= rows; i++ {
	// 	//inner loop for spaces before stars
	// 	for j := 1; j <= rows-i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	//inner loop for stars
	// 	for k := 1; k <= 2*i-1; k++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println() // Move to next line
	// }

	// for i := range 10 {
	// 	fmt.Println(10 - i)
	// }

	// i := 5
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	//generate a random number between 1 and 100
	target := random.Intn(100) + 1

	// Welcome message
	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I have chosen a number between 1 and 100")
	fmt.Println("Can you guess what it is ?")

	var guess int
	for {
		fmt.Println("Enter your Guess")
		fmt.Scanln(&guess)

		// Check if the guess is correct
		if guess == target {
			fmt.Println("Congratulations You Guessed correct number!")
			break
		} else if guess < target {
			fmt.Println("Too low! Try guessing a higher number")
		} else {
			fmt.Println("Too high! Try guessing a lower number")
		}
	}

}
