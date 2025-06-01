package main

import "fmt"

func main() {
	// Switch statement in go is (switch case default) (fallthrough)
	// switch expression {
	// case value1:
	// 	// code to be executed if statement equals to value1
	//  // fallthrough
	// case value2:
	// 	// code to be executed if statement equals to value2
	// case value3:
	// 	// code to be executed if statement equals to value3
	// default:
	// // Code to be ececutes if expression does not match any value
	// }

	// Switch statement
	// switch expression {
	// case value1:
	// 	// code to be executed if statement equals to value1
	//  // break;
	// case value2:
	// 	// code to be executed if statement equals to value2
	//  // break;
	// case value3:
	// 	// code to be executed if statement equals to value3
	// default:
	// // Code to be ececutes if expression does not match any value
	// // break;
	// }

	fruit := "apple"
	switch fruit {
	case "apple":
		fmt.Println("It is an apple.")
	case "banana":
		fmt.Println("It is a banana.")
	default:
		fmt.Println("Unknown Fruit!")
	}

	// Multiple Conditions
	day := "Monday"

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thrusday", "Friday":
		fmt.Println("It's a weekday")
	case "Saturday", "Sunday":
		fmt.Println("It's a weekend")
	default:
		fmt.Println("Invalid day.")
	}

	number := 15
	switch {
	case number < 10:
		fmt.Println("Number is less then 10")
	case number >= 10 && number < 20:
		fmt.Println("Number is between 10 and 19")
	default:
		fmt.Println("Number is 20 or more")
	}

	num := 2
	switch {
	case num > 1:
		fmt.Println("Greater then 1")
		fallthrough
	case num == 1:
		fmt.Println("Number is 2")
	default:
		fmt.Println("Not Two")
	}

	checkType(10)
	checkType(3.14)
	checkType("Hello")
	checkType(true)
}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("It's an integer")
	case float64:
		fmt.Println("It's float")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("Unknown type")
	}
}
