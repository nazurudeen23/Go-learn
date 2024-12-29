package main

import (
    "fmt"
)

// Function to demonstrate arithmetic operations
func arithmeticOperations() {
    a, b := 10, 5
    fmt.Println("Arithmetic Operations:")
    fmt.Printf("Addition: %d + %d = %d\n", a, b, a+b)
    fmt.Printf("Subtraction: %d - %d = %d\n", a, b, a-b)
    fmt.Printf("Multiplication: %d * %d = %d\n", a, b, a*b)
    fmt.Printf("Division: %d / %d = %d\n", a, b, a/b)
    fmt.Printf("Modulus: %d %% %d = %d\n", a, b, a%b)
}

// Function to demonstrate relational operations
func relationalOperations() {
    x, y := 10, 20
    fmt.Println("\nRelational Operations:")
    fmt.Printf("%d == %d: %v\n", x, y, x == y)
    fmt.Printf("%d != %d: %v\n", x, y, x != y)
    fmt.Printf("%d > %d: %v\n", x, y, x > y)
    fmt.Printf("%d < %d: %v\n", x, y, x < y)
    fmt.Printf("%d >= %d: %v\n", x, y, x >= y)
    fmt.Printf("%d <= %d: %v\n", x, y, x <= y)
}

// Function to demonstrate logical operations
func logicalOperations() {
    a, b := true, false
    fmt.Println("\nLogical Operations:")
    fmt.Printf("true && false: %v\n", a && b)
    fmt.Printf("true || false: %v\n", a || b)
    fmt.Printf("!true: %v\n", !a)
}

// Function to demonstrate bitwise operations
func bitwiseOperations() {
    p, q := 5, 3 // In binary: 5 = 0101 and 3 = 0011
    fmt.Println("\nBitwise Operations:")
    fmt.Printf("Bitwise AND: %d & %d = %b\n", p, q, p&q) // 0101 & 0011 = 0001 (1)
    fmt.Printf("Bitwise OR: %d | %d = %b\n", p, q, p|q)   // 0101 | 0011 = 0111 (7)
    fmt.Printf("Bitwise XOR: %d ^ %d = %b\n", p, q, p^q) // 0101 ^ 0011 = 0110 (6)
    fmt.Printf("Bitwise NOT: ^%d = %b\n", p, ^p)         // NOT operation
}

// Function to demonstrate assignment operations
func assignmentOperations() {
    var c int = 10
    c += 5 // Add and assign
    fmt.Println("\nAssignment Operations:")
    fmt.Println("After c += 5:", c) // c is now 15

    c -= 3 // Subtract and assign
    fmt.Println("After c -= 3:", c) // c is now 12

    c *= 2 // Multiply and assign
    fmt.Println("After c *= 2:", c) // c is now 24

    c /= 4 // Divide and assign
    fmt.Println("After c /= 4:", c) // c is now 6

    c %= 5 // Modulus and assign
    fmt.Println("After c %= 5:", c) // c is now 1
}

func main() {
    arithmeticOperations()
    relationalOperations()
    logicalOperations()
    bitwiseOperations()
    assignmentOperations()
}
