package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println(m)

	m["answer"] = 23
	fmt.Println("value is: ", m["answer"])

	m["answer"] = 97
	elem := m["answer"]
	fmt.Println("value is: ", elem)

	// Delete an element
	delete(m, "answer")
	fmt.Println("value is: ", m)

	v, ok := m["answer"]
	fmt.Println("value: ", v, "Present? ", ok)

}
