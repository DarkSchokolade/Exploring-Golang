package main

import "fmt"

func main() {
	sum := 1
	//	The init and post statements are optional.
	// acts as a while loop [!There is no while loop in go]
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
}

// for{} is a forever loop
