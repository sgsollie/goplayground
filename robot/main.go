// Robot places a virtual robot on an imaginary unlimited grid and moves it
package main

import (
	"fmt"
	"log"
)

// Positon will give the robots current coordinates
type position struct {
	xpos int
	ypos int
}

// vector is actually vectors and heading
type vector struct {
	heading  string
	location position
}

// turn takes a heading as d and a vector value, and sets the heading of the robot
func Turn(d string, v *vector) *vector {
	switch d {
	case "north":
		v.heading = "north"
	case "south":
		v.heading = "south"
	case "east":
		v.heading = "east"
	case "west":
		v.heading = "west"
	default:
		log.Fatalln("Invalid direction entered: ", d, " Valid directions are: north, south, east, west")
	}
	return v
}

// Advance moves the robot forward along a given axis
func Advance(v *vector) *vector {
	switch v.heading {
	case "north":
		v.location.xpos += 1
	case "south":
		v.location.xpos += 1
	case "east":
		v.location.ypos += 1
	case "west":
		v.location.ypos += 1
	}
	return v
}

func main() {
	//Place a robot in an initial position on an infinite grid
	robot1 := vector{
		heading:  "north",
		location: position{2, 2},
	}
	fmt.Println(robot1)

	//Move robot1 forward two steps
	//Turn east
	//Move forward another step
	//Print new location and heading
	Advance(&robot1)
	Advance(&robot1)
	Turn("east", &robot1)
	Advance(&robot1)
	fmt.Println(robot1)
}
