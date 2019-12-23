package main

import "fmt"

func main() {

	ts := filternumbers([]int{10, 20, 30, 40}, func(n int) bool {
		return n > 10
	})

	fmt.Println(ts)
}

func filternumbers(numbers []int, callback func(int) bool) []int {
	ts := []int{}
	for _, n := range numbers {
		if callback(n) {
			ts = append(ts, n)
		}
	}
	return ts
}
