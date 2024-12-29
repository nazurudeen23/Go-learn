package main

import "fmt"

func main() {
	forLoop()
	whileLoop()
	rangeLoop()
}


func forLoop(){
	// For loop
	for i := 1; i <= 5; i++ {
		fmt.Println("Iteration:", i)
	}
}

func whileLoop(){
	// While loop
	num := 1
	for num <= 5 {
		fmt.Println("While loop Iteration:", num)
		num++
	}
}

func rangeLoop(){
	// Range loop
	fruits := []string{"apple", "banana", "cherry"}
    for index, fruit := range fruits {
        fmt.Printf("Fruit at index %d: %s\n", index, fruit)
    }
}