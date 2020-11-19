package main

import "fmt"

func main() {
	sum := 0
	// do not use brackets in for loop it'll throw error
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
