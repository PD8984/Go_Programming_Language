package main

import "fmt"

func main() {
	var a int = 69 // Declares an integer variable 'a' and initializes it with the value 69.
	// 'a' now holds the value 69 directly in its memory location.

	var aptr *int = &a // Declares a pointer variable 'aptr'.
	// The '*' in '*int' means 'aptr' is a pointer to an integer.
	// The '&' (address-of operator) gets the memory address of 'a'.
	// So, 'aptr' now stores the memory address where the value of 'a' (69) is located.

	*aptr = *aptr + 1 // This line performs two key operations using the '*' (dereference operator):
	// 1. '*aptr' on the right side: It accesses the value currently stored at the memory address held by 'aptr' (which is the value of 'a', i.e., 69).
	// 2. `... + 1`: It adds 1 to that value (69 + 1 = 70).
	// 3. '*aptr' on the left side: It writes the new calculated value (70) back to the memory address held by 'aptr'.
	// Essentially, this line directly modifies the value of 'a' through the pointer 'aptr'. 'a' now becomes 70.

	fmt.Println(&a) // Prints the **memory address** of the variable 'a'.
	// This will be a hexadecimal number, like 0xc0000120a0.

	fmt.Println(a) // Prints the **current value** of the variable 'a'.
	// After the modification through 'aptr', this will print 70.

	fmt.Println(aptr) // Prints the **value stored inside the pointer variable 'aptr'**.
	// This value is the memory address that 'aptr' points to, which is the same as &a.

	fmt.Println(*aptr) // Prints the **value pointed to by 'aptr'**.
	// The '*' (dereference operator) accesses the data at the address 'aptr' holds.
	// This will also print 70, as it's the current value of 'a'.

	fmt.Println(&aptr) // Prints the **memory address of the pointer variable 'aptr' itself**.
	// 'aptr' is a variable, and like any variable, it resides at its own memory address.
	// This will be a different memory address from &a.

}

