package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	// Here we are using log.Fatalln() to exit the program when an error occurs.
	// Fatalln quits the program so we don't need to specify "return"
	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
// here we are returning both the error using fmt.Errorf()
// and an empty slice of byte as there are two return types specified
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("Error found: %v", err)
	}
	return bs, nil
}
