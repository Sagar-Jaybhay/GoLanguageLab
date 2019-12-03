package main

import (
	"fmt"
)

func main() {

	var i = 1

	if 1 == 1 {
		if i > 9 {
			fmt.Println("i m in if block")
		} else if i < 1 {
			fmt.Println("i am in else if block")
		} else {
			fmt.Println("i am in else block")
		}

	} else {
		fmt.Println("variable i not equal to 1")
	}

}
