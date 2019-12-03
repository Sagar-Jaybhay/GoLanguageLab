package main

var y = "sagar jaybhay"

func main() {

	//x := 10
	//fmt.Println("Program working correctly ", x)

	///fmt.Println("i am accessing outer variable inside function := ", y)

	variable := func() int {
		return (10 + 20)
	}
	println(variable())

}
// println(variable())  /// not able to call this function outside of scope
//fmt.Println(x)
