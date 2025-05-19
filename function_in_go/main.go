package main

import "fmt"

// Function to add two integers
func add(a int, b int) int {
    return a + b
}

// Function to subtract two integers
func subtract(a int, b int) int {
    return a - b
}

// Function to multiply two integers
func multiply(a int, b int) int {
    return a * b
}

// Function to divide two integers
func divide(a int, b int) float64 {
    if b == 0 {
        fmt.Println("Error: Division by zero")
        return 0
    }
    return float64(a) / float64(b)
}


func main() {
    // Calling functions and displaying results
    sum := add(5, 10)
    fmt.Println("Addition:", sum) // Output: Addition: 15

    difference := subtract(10, 5)
    fmt.Println("Subtraction:", difference) // Output: Subtraction: 5

    product := multiply(4, 3)
    fmt.Println("Multiplication:", product) // Output: Multiplication: 12

    quotient := divide(20, 4)
    fmt.Println("Division:", quotient) // Output: Division: 5

    // Attempting division by zero
    divide(10, 0)
}
