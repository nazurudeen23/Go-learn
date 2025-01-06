package main

import (
	"fmt"
	"math"
)

type Shape interface{
	Area() float64
}

//struct
type Rectangle struct{
	width, height float64
}


//method implementation
func (r Rectangle) Area() float64{
  return r.width * r.height
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
func calculateArea(s Shape) float64{
    return s.Area()
}

func main(){
	rect := Rectangle{width: 5, height: 10}
	circle := Circle{radius: 2}

	fmt.Println("Rectangle area: ", calculateArea(rect))
	fmt.Println("Circle area: ", calculateArea(circle))
}