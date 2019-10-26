//An imaginary virtual guitar tuner, using METHODS to "tune" a vitual guitar 
// with (values g1 and g2)

package main

import (
	"fmt"
)

//Setting our user-defined type "guitar"
type guitar struct {
	make     string
	model    string
	electric bool
	tuning   []string
}

//METHOD (or special function with receiver arguments, used for custom types)
//func (r receivertype) identifier(args) returntype { code }
func (r guitar) Etuner() []string {
	r.tuning = []string{`E`, `A`, `D`, `G`, `B`, `E`}
	return r.tuning
}

func (r guitar) Dtuner() []string {
	r.tuning = []string{`D`, `A`, `D`, `G`, `B`, `E`}
	return r.tuning
}

func main() {
	//Add a VALUE for TYPE guitar "g1" and "g2"
	g1 := guitar{
		`Ibanez`,
		`ARX`,
		true,
		[]string{},
	}

	g2 := guitar{
		`Fender`,
		`Strat`,
		true,
		[]string{},
	}
	fmt.Println(g1)

	//Put g1 "out of tune" initially, so we can "tune" it, with one of our functions
	g1.tuning = []string{`F#`, `G#`, `D`, `G`, `B`, `Eb`}
	fmt.Println(g1)

	//Try Tuning it using the Etuner function: - applying the method while printing
	fmt.Println(g1.Etuner())

	//Doing the same with g2, but updating the VALUE, then printing
	g2.tuning = g2.Dtuner()
	fmt.Println(g2)
}

