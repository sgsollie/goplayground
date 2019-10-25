// An exercise with maps, containing []string types as values

package main

import (
	"fmt"
)

func main() {

	x1 := []string{`Shaken, not stirred`, `Martinis`, `Women`}
	x2 := []string{`James Bond`, `Literature`, `Computer Science`}
	x3 := []string{`Being evil`, `Ice cream`, `Sunsets`}

	x := map[string][]string{
		`bond_james`:      x1,
		`moneypenny_miss`: x2,
		`no_dr`:           x3,
	}
	//fmt.Println(x)

	x[`Ollie`] = []string{`Flying`, `Coding`, `Skiing`}

	delete(x, `no_dr`)

	for k, v := range x {
		fmt.Printf("%v\n", k)
		for i, w := range v {
			fmt.Printf("\t%v, %v\n", i, w)
		}
	}

}
