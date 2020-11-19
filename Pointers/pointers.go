package main

import "fmt"

func main() {
	i, j := 42, 127

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
	fmt.Println(j)
}
