package main

import "fmt"

func main() {
	defer hellowword()
	hello()

	//fmt.Println("i m here for defer keyword")
}

func hello() {
	fmt.Println("i am in hello")
}

func hellowword() {
	fmt.Println("i am in hello world")
}
