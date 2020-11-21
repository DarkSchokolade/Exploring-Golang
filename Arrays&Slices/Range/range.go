package main

import "fmt"

func main() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// range continued
	// range returns 2 output index and value
	// use _ [underscore] to omit either
	pow2 := make([]int, 5)
	for i := range pow2 {
		pow2[i] = 1 << uint(i) // bit shift left
	}

	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}
}
