package main

import "fmt"

func main() {

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Printf("%v %T", primes)
	fmt.Println("Printing complete slice", primes)
	var s []int = primes[1:4] // this is slicing the slice
	fmt.Println(primes[2])    //index access ; access by index
	fmt.Println(s)

	var myslice []string
	fmt.Print(myslice)
	fmt.Printf()

	//------------------ appending elements to slice beyond it capacity -----------------------------
	myslice := make([]string, 3, 5)
	// 3 is length:- it is number of elements refered by slice
	// 5 is capacity:- it is number of elements in underlying array.

	myslice[0] = "Sagar"
	myslice[1] = "Jaybhay"
	myslice[2] = "Sr. Software Developer"

	fmt.Println(myslice)
	fmt.Println("Orignal Length ", len(myslice))
	fmt.Println("Orignal Capcity ", cap(myslice))

	myslice = append(myslice, "Pune")
	myslice = append(myslice, "Maharashtra")
	myslice = append(myslice, "India")

	fmt.Println("Updated ", myslice)
	fmt.Println("Updated Length ", len(myslice))
	fmt.Println("Updated Capcity ", cap(myslice))

	// ----------- calling different ways to create the slice ---------------------

	differentWaystocreateSlice()

	//-----------call different ways to create slice
	shorthandsyntax()

	//

	declarewithvar()

	withmake()
}

func differentWaystocreateSlice() {

	records := make([][]string, 2)

	emp := make([]string, 2)
	emp[0] = "A"
	emp[1] = "B"

	records = append(records, emp)
	fmt.Println("first array appended")
	fmt.Println(records)

	emp1 := make([]string, 2)
	emp1[0] = "A"
	emp1[1] = "B"

	records = append(records, emp1)

	fmt.Println("Second array appended")
	fmt.Println(records)

}

func shorthandsyntax() {
	record := []string{}

	fmt.Println(record)        //o/p = []
	fmt.Println(record == nil) // o/p  = false

}

func declarewithvar() {

	var record []string
	fmt.Println(record)        //o/p = []
	fmt.Println(record == nil) // o/p  = true

}

func withmake() {
	record := make([]string, 20)
	fmt.Println(record)        //o/p = []
	fmt.Println(record == nil) // o/p  = false

}
