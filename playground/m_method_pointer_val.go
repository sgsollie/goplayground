package main

import (
	"fmt"
)

type dog struct {
	name  string
	breed string
	age   int
	toys  []string
}

// method to give dog a toy
// this works by passing an unlimited amount of strings as an argument (identified as toy)
// this is then added as a slice of string to the .toys field of the struct
// we have to use pointer to type dog as this is the correct way to update the value, 
// if we don't use a pointer, the value doesn't get updated as hoped, it just duplicates the value within the method 
//& doesn't update the one referenced in main()

func (d *dog) giveToy(toy ...string) {
	//fmt.Println(toy)
	(*d).toys = []string(toy)
}

func main() {
	lincoln := dog{
		name:  `Lincolnshire`,
		breed: `Sausage`,
		age:   0,
	}

	//calling the method: 
	lincoln.giveToy("Ball", "Tug", "Bone")
	fmt.Println(lincoln)
}

