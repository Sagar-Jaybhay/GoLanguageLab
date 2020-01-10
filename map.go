package main

import "fmt"

func main() {

	m := make(map[string]int) // create map here
	m["k1"] = 7
	m["k2"] = 10

	fmt.Println(m) //map[k1:7 k2:10]

	delete(m, "k1")
	fmt.Println(m) //map[k2:10]

	pt, kt := m["k2"]

	fmt.Println(pt, kt) //10 true

}
