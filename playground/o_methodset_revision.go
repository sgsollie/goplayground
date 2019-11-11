package main

import "fmt"

// User defined structs
type person struct {
	first string
	last  string
	age   int
}

// Interface named human, implementing the speak method, which means the reciever type
// of these methods are also of type "human" now in a sense. (or more correctly, those
// methods implement the "human" interface)
type human interface {
	speak()
}

//The speak() method taking a pointer to type person
func (r *person) speak() {
	fmt.Println((*r).first, "is saying hello!")
}

//saySomething() func, implementing the human interface (and as a result, pointer to a value of type person)
func saySomething(v human) {
	v.speak()
}

func main() {
	p1 := person{
		first: "Oliver",
		last:  "Young",
		age:   31,
	}

	//Works:
	saySomething(&p1)

	//Doesnt' Work, comment the below line out to run this program
	//saySomething(p1)
}
