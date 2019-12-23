package main

import "fmt"

func main() {

	var x [10]int

	fmt.Println("Printing array : ", x)

	fmt.Println("array length : ", len(x))

	fmt.Println("x[4] value is  : ", x[4])

	for i := 0; i < 10; i++ {
		x[i] = 100 + i*20
	}
	fmt.Println("printing added value array :")
	fmt.Println(x)

}
