package main

import (
	"fmt"
	"maps"
)

func main() {
	// var m map[keyType]valueType

	// mapVariable = make(map[keytype]valueType)

	// using a Map Litral
	// mapVariable = map[keyType]valueType{
	// 	key1 : value1,
	// 	key2 : value2
	// }

	myMap := make(map[string]int)
	fmt.Println(myMap)
	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3
	fmt.Println(myMap)

	delete(myMap, "two")
	fmt.Println(myMap)

	// clear(myMap)
	// fmt.Println(myMap)

	value, unknownvalue := myMap["one"]
	fmt.Println(value)
	fmt.Println(unknownvalue)

	myMap2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(myMap2)

	if maps.Equal(myMap, myMap2) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	for k, v := range myMap2 {
		fmt.Println(k, v)
	}
}
