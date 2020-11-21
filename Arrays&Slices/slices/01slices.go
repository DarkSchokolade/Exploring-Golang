package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13} //array
	var s []int = primes[1:4]            //this is a slice [half open range]
	fmt.Println(s)

	s[0] = 99 // slice is a flexible view into the array and modifies the array
	fmt.Println(primes)
	// A slice does not store any data, it just describes a section of an underlying array.

	// although copying an array and changing does not affect each other
}
