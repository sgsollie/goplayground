//Unmarshalling JSON

package main

import (
	"encoding/json"
	"fmt"
)

//Define a type named person, with underlying type of struct:
//Including json fields

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {

	//Json blob defined as s
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`

	//Defining target variable bs as type slice of person
	var bs []person

	//unmarshal the json, which takes the blob as the first argument (as a slice of byte, hence the conversion []byte(s) and then stores it at &bs
	err := json.Unmarshal([]byte(s), &bs)

	//print it to check it stored it, and do error handling if the json happens to be invalid/corrupt.
	fmt.Println(bs)
	fmt.Printf("Type of unmarshalled json is: %T\n\n", bs)
	if err != nil {
		fmt.Println("error:", err)

	}

	//now the json is stored as a value of type []person we can use it in our program
	// here we are ranging over the values and printing them:
	for k, v := range bs {
		fmt.Printf("%v\n", k)
		fmt.Printf("\t%v %v %v Sayings:\n", v.First, v.Last, v.Age)
		for _, v := range v.Sayings {
			fmt.Printf("\t\t%v\n", v)
		}
	}
}

