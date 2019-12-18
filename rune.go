package main

func main() {

	for i := 50; i < 150; i++ {
		println(i, " --- ", string(i), " --- ", []byte(string(i)))
	}

	// foo := 'a'
	// println(foo)
	// fmt.Printf("%T", foo)
	// println()
	// fmt.Printf("%T", rune(foo))

}
