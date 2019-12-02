package utilitypackage

import (
	"fmt"
)

// this name field is used in main package
var Name string = "sagar"

// This CreateDataTypes called in main function
func CreateDataTypes() {
	a := 10
	b := true
	c := 12.144
	d := "Sagar Jaybhay"
	e := complex(3.7, 4.8)

	fmt.Printf("%v \t", a)
	fmt.Printf("%v \t ", b)
	fmt.Printf("%v \t", c)
	fmt.Printf("%v \t", d)
	fmt.Printf("%v \t", e)

	fmt.Println("\n-------------------------------------------------")

	fmt.Printf("%T \t", a)
	fmt.Printf("%T \t ", b)
	fmt.Printf("%T \t", c)
	fmt.Printf("%T \t", d)
	fmt.Printf("%T \t", e)
}
