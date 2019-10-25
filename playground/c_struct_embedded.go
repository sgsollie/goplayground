//Example of embedded structs

package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	//Have to assign values to vehicle too as it is embedded
	t := truck{
		vehicle: vehicle{
			doors: 5,
			color: "blue",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 3,
			color: "black",
		},
		luxury: true,
	}

	fmt.Println(t, s)

	//Though vehicle struct is embedded, when calling values it gets promoted,
	//so we can just use struct name t and call color directly:
	fmt.Println(t.color)

	//And another embedded value:
	fmt.Println(s.luxury)

}

