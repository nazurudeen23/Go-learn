package main

import "fmt"

type Employee struct {	// Defining a struct
	name string
	age int
}

func main() {
	boolean()
	integerType()
	floatType()
	stringType()
	complexNumber()
	array()
	slice()

	// Declaring a struct variable
	var emp Employee
	emp.name = "Alice"
	emp.age = 25
	fmt.Println("Employee Name:", emp.name, "Age:", emp.age)
}


func boolean(){
	// Declaring a boolean variable
	var isTrue bool = true
	fmt.Println(isTrue)
}

func integerType(){
	var a int = 42          // Signed integer
    var b uint = 100       // Unsigned integer
    fmt.Println("Signed:", a) // Output: Signed: 42
    fmt.Println("Unsigned:", b) // Output: Unsigned: 100
}

func floatType(){
	var a float64 = 42.5
	b := 42.5
	fmt.Println("Value of a is", a) // Output: Value of a is 42.5
	fmt.Println("Value of b is", b) // Output: Value of b is 42.5
}

func stringType(){
	// Declaring a string variable
	var name string = "Alice"
	fmt.Println(name)
}

func complexNumber(){
	// Declaring a complex number
	var a complex64 = 5 + 5i
	fmt.Println(a)
}

func array(){
	// Declaring an array
	var numbers [4]int = [4]int{1, 2, 3, 4}
	fmt.Println(numbers)
}

func slice(){
	// Declaring a slice
	var numbers = []int{1, 2, 3, 4}
	fmt.Println(numbers)
}