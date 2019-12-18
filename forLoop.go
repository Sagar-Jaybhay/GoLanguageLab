package main

func main() {
	println("I am In Main Function")

	// for i := 0; i < 10; i++ {
	// 	println("value of i ", i)
	// }

	i := 0

	// for i < 10 { // another type of for loop
	// 	println("value of i: ", i)
	// 	i++
	// }

	for {

		i++

		if i%2 == 0 {
			continue
		}
		println("val of i :=> ", i)

		if i >= 30 {
			break
		}

	}

}
