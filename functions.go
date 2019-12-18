package main

import "fmt"

func displaycustomer(name string) {
	fmt.Println(name)
}

func main() {
	fmt.Println("i am in main function")

	displaycustomer("sagar")
	displaycustomer("jaybhay")

	b := addition(10, 20)
	fmt.Println("addition is => ", b)

	//s := fullname("sagar", "jaybhay")
	print(returnnamedvalue("sagar", "jaybhay"))
	println()
	fmt.Print(multiplevaluesreturn(10, 20))
}

func addition(a, b int) int {
	return a + b
}

func fullname(fname string, lname string) string {
	return fmt.Sprint(fname, lname)
}

func returnnamedvalue(fname, lname string) (s string) {
	return fmt.Sprint(fname, lname)
}

func multiplevaluesreturn(a, b int) (int, int) {
	return a * 10, b * 20
}
