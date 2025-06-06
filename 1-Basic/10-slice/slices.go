package main

import (
	"fmt"
	"slices"
)

func main() {
	// var slice []ElementType

	// var number []int
	// var numbers1 = []int{1, 2, 3}
	// number2 := []int{9, 8, 7}

	slice := make([]int, 5)
	a := [5]int{1, 2, 3, 4, 5}
	slice1 := a[1:4]

	fmt.Println(slice)
	fmt.Println(slice1)

	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1)
	fmt.Println(sliceCopy)

	var nilSlice []int
	fmt.Println(nilSlice)

	for i, v := range slice1 {
		fmt.Println(i, v)
	}

	if slices.Equal(slice1, sliceCopy) {
		fmt.Println("Slice1 is equal to sliceCopy")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
			fmt.Printf("adding value %d in outer slice at index %d, and in inner slice index of %d\n", i+j, i, j)
		}
	}
	fmt.Println(twoD)

	slice2 := slice1[2:4]
	fmt.Println(slice2)

	fmt.Println("The capacity of slice2 is", cap(slice2))
	fmt.Println("The len of slice2 is", len(slice2))
}
