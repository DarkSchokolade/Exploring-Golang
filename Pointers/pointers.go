package main

import "fmt"

func main() {
	i, j := 42, 120

	p := &i         // points to i
	fmt.Println(*p) // read i through pointer
	*p = 21         // set i through pointer
	fmt.Println(i)
	fmt.Println("address of i: ", &p) //address of pointer
	p = &j                            // point to j
	*p = *p / 3                       //divide j through pointer
	fmt.Println(j)
}
