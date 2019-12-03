package main

import "fmt"

func main() {
	var a = 2

	switch a {
	case 1:
		fmt.Printf("1")
	case 2:
		fmt.Printf("2")
	case 3:
		fmt.Printf("3")
	}
	anotherswitch()
}

func anotherswitch() {

	println("i am in another_switch function")
	var number = 40

	switch {
	case number == 10:
		println("number is 10")
		number = number + 10

	case number == 20:
		println("number is 20")

	default:
		println("nothing has found")

	}
	println("number is ", number)
}
