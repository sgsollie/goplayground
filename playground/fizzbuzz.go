//The sole parameter for my function will be an integer.
//The sole output of my function will be an incrementing array, starting from the number 1 and ending at the input number.
//If a number is a multiple of 3, the output will instead be `fizz`.
//If a number is a multiple of 5, the output will instead be `buzz`.
//However, if the output is a multiple of both 3 and 5, the output will instead be `fizzbuzz`.

//Taken from: https://hackernoon.com/the-best-whiteboard-interview-advice-i-ever-received-3ebbfa72e4a

package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(num int) []string {
	var fizzval string
	fizzslice := make([]string, num+1)
	for i := 1; i <= num; i++ {

		if i%3 == 0 && i%5 == 0 {
			//fmt.Println("FizzBuzz")
			fizzval = "FizzBuzz"
		} else if i%5 == 0 {
			//fmt.Println("Buzz")
			fizzval = "Buzz"
		} else if i%3 == 0 {
			//fmt.Println("Fizz")
			fizzval = "Fizz"
		} else {
			fizzval = strconv.Itoa(i)
		}

		fizzslice[i] = fizzval

	}

	return fizzslice
}

func main() {
	x := fizzbuzz(20)
	//fmt.Println(x)

	for k, v := range x {
		fmt.Printf("%v: %v\n", k, v)
	}

}
