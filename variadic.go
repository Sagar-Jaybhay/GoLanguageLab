package main

import "fmt"

func main() {
	fmt.Print("This is variadic function demo")
	n := totalsum(10, 20, 30, 40, 50)
	fmt.Println("total sum is ", n)

}

func totalsum(array ...int) int {

	fmt.Println(array)
	fmt.Println("------------------")

	fmt.Printf("%T \t", array)

	total := 0

	for _, v := range array {
		fmt.Print(_)
		println(v)
		total += v
	}
	return total
}
