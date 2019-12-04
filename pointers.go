package main

func retunbyreference(x *int) {
	*x = (*x) * 10
	println("value of x in retunbyreference method", *x,
		"  memory address of x in retunbyreference method ", &(*x))

}

func main() {

	x := 10
	retunbyreference(&x)
	println("value of x in main method", x,
		"  memory address of x in main method ", &x)

}

func returnbyvalue(x int) {
	x = x * 10
	println("value of x in returnbyvalue method", x,
		"  memory address of x in returnbyvalue method ", &x)
}
