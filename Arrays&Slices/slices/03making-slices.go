package main

import "fmt"

// Slices can be created with make function
// This is how you create dynamically sized arrays
func main() {
	a := make([]int, 5) // Both length and capacity are set to 5
	printSlice("a", a)

	b := make([]int, 0, 5) // To specify capacity pass 3rd argument
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	// Extra
	var s []int // Dynamically sized
	s = append(s, 3, 4, 5)
	printSlice("s", s)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d capacity=%d %v\n", s, len(x), cap(x), x)
}
