package main

import "fmt"

func main() {
	s := newSlice()

	s.print()
}

type sliceInt []int

func newSlice() sliceInt {
	return []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
}

func (s sliceInt) print() {
	for _, number := range s {

		var value string = "odd"
		if number%2 == 0 {
			value = "even"
		}
		fmt.Println(number, value)
	}
}
