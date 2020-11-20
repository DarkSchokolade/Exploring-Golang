package main

import "fmt"

// arrays cannot be resized
// this seems limiting but checkout slices
func main() {
	var a [3]string
	a[0] = "bob"
	a[1] = "the builder"
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
