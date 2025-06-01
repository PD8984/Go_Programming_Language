package main

import "fmt"

func main() {
	// constants holds values like variables but are immutable in nature
	// Define constants for weekdays using iota for sequential values.
	// iota starts at 0, so Monday will be 1, Tuesday 2, and so on.
	const (
		Monday    = iota + 1 // 0
		Tuesday              // 1
		Wednesday            // 2
		Thursday             // 3
		Friday               // 4
		Saturday             // 5
		Sunday               // 6
	)

	fmt.Println("Weekday Constants")
	fmt.Println("Monday:", Monday)
	fmt.Println("Tuesday:", Tuesday)
	fmt.Println("Wednesday:", Wednesday)
	fmt.Println("Thursday:", Thursday)
	fmt.Println("Friday:", Friday)
	fmt.Println("Saturday:", Saturday)
	fmt.Println("Sunday:", Sunday)
}
