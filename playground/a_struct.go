//A little playing around with a struct.

package main

import (
	"fmt"
)

type aircraft struct {
	name       string
	powerplant string
}

func NewSEPAircraft(name string) *aircraft {

	a := aircraft{name: name}
	a.powerplant = `Propeller`
	return &a
}

func main() {
	s := aircraft{name: "Extra300", powerplant: `Propeller`}
	fmt.Println(s.powerplant)

	v := NewSEPAircraft(`Cessna150`)
	fmt.Println(v)

}

