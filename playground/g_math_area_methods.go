// Created methods to deal with calculating the area of a square and a circle.

package main

import (
	"fmt"
	"math"
)

const Pi float64 = 3.14159265358979323846264338327950288419716939937510582097494459

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (r square) squareArea() float64 {
	area := r.length * r.width
	return area
}

func (r circle) circleArea() float64 {
	//area := Pi * r.radius * r.radius
	area := Pi * math.Pow(r.radius, 2)
	return area
}

func main() {

	//Printing the area of a square, using the squareArea() method for type square
	sq1 := square{2.5, 2.5}
	fmt.Println(sq1.squareArea())

	//Printing the area of a circle, using Pi and the circleArea() method
	circ1 := circle{3.9}
	fmt.Println(circ1.circleArea())
}

