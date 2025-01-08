package main

import (
	"fmt"
	"math"
)

// Shape Interface
type Shape interface {
	Area() float64
}

// Measurable Interface
type Measurable interface {
	Perimeter() float64
}

// Geometry Interface
type Geometry interface {
	Measurable
	Shape
}

// Rectangle Struct
type Rectangle struct {
	width, height float64
}

// Area Method for Rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Perimeter Method for Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Circle Struct
type Circle struct {
	radius float64
}

// Area Method for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Perimeter Method for Circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Describe Shape Function
func discribeShape(g Geometry) {
	fmt.Println("Area: ", g.Area())
	fmt.Println("Perimeter: ", g.Perimeter())
}

// CalculationError Struct
type CalculationError struct {
	msg string
}

// Error Method for CalculationError
func (ce CalculationError) Error() string {
	return ce.msg
}

// Perform Calculation Function
func performCalculation(val float64) (float64, error) {
	if val < 0 {
		return 0, CalculationError{
			msg: "Invalid Input",
		}
	}
	return math.Sqrt(val), nil
}

func main() {
	rect := Rectangle{width: 5, height: 10}
	discribeShape(rect)

	// Handle Calculation
	val, err := performCalculation(rect.height)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Value:", val)
	}
}
