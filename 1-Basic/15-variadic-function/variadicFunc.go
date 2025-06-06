package main

import "fmt"

func main() {
	// // ... Ellipsis
	// func functionName(param1 ...type1 returnType {
	// 	//function body
	// }

	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 23, 34, 45, 65, 67, 7889, 90))

}
func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}
