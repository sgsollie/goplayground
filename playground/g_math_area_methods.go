// Created methods to deal with calculating the area of a square and a circle.

package main

import (
	"fmt"
	"math"
)

// we obviously know pi, and that it won't change:
const Pi float64 = 3.14159265358979323846264338327950288419716939937510582097494459

//user defined types for Square and circle
type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

//Interface for shape methods, which will take all methods of Area (and as a result, types of square and circle)
type shape interface {
	Area() float64
}

//Area method for square type
func (r square) Area() float64 {
	area := r.length * r.width
	return area
}

//Area method for circle type
func (r circle) Area() float64 {
	//area := Pi * r.radius * r.radius
	area := Pi * math.Pow(r.radius, 2)
	return area
}

func info(num shape) {
	fmt.Println(num.Area())
}

func main() {

	//Printing the area of a square, using the Area() method for type square
	sq1 := square{2.5, 2.5}
	fmt.Println(sq1.Area())

	//Printing the area of a circle, using Pi and the Area() method
	circ1 := circle{3.9}
	fmt.Println(circ1.Area())

	//Printing the area using the shape interface and info function.
	info(circ1)

}

