package main

import "fmt"

func main() {
    var name string = "Alice" // Declare and initialize a variable
    age := 25                 // Short declaration
    
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    
    // Changing the value of the variable
    age = 30
    fmt.Println("Updated Age:", age)

	const MaxUsers = 100 // Declare a constant
    
    fmt.Println("Maximum Users Allowed:", MaxUsers)
}
