package main

import (
	"fmt"
)


// Reverse takes a string and prints it backwards
func Reverse(s string) {
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Printf("%c", s[i])
	}

}

func main() {
	Reverse("Oliver")
}

