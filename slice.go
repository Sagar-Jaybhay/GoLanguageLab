package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	//fmt.Printf("%v %T", primes)
	fmt.Println("Printing complete slice", primes)
	var s []int = primes[1:4] // this is slicing the slice
	fmt.Println(primes[2])    //index access ; access by index

	fmt.Println(s)
}
