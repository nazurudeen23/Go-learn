package main

import "fmt"

func main() {
	// Declaring an integer variable
	var a int

	// Declaring a pointer variable
	var ptr *int

	// Initializing the pointer variable with the address of variable a
	ptr = &a

	// Assigning a value to a variable
	a = 10

	// Accessing the value stored in the address
	// pointed by the pointer variable
	fmt.Println("Value stored in variable a:", a) // Output: Value stored in variable a: 10
	fmt.Println("Address of variable a:", &a) // Output: Address of variable a: 0xc0000140b0
	fmt.Println("Value stored in variable a:", *ptr) // Output: Value stored in variable a: 10
	fmt.Println("Address stored in pointer variable:", ptr) // Output: Address stored in pointer variable: 0xc0000140b0
	fmt.Println("Value stored at the address stored in pointer variable:", &ptr) // Output: Value stored at the address stored in pointer variable: 10
}