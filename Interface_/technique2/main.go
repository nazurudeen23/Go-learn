package main

import (
	"fmt"
	"math"
)

type Shape interface{
	Area() float64
}

//----------------------------------------
type Measurable interface{
   Perimeter() float64
}

type Geometry interface{
	Measurable
	Shape
} 
//----------------------------------------

//struct
type Rectangle struct{
	width, height float64
}

//method implementation for Area
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

//method implementation
func (r Rectangle) Perimeter() float64{
  return 2*(r.width * r.height)
}

//struct
type Circle struct{
	radius float64
}

//method implementation
func (c Circle) Area() float64{
	return math.Pi * c.radius * c.radius
}

//function 
func discribeShape(g Geometry) {
    fmt.Println("Area: ", g.Area())
	fmt.Println("Perimeter: ",g.Perimeter())
}

func main(){
	rect := Rectangle{width: 5, height: 10}
	discribeShape(rect)
	

}
