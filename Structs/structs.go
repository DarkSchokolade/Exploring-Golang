package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v1 := Vertex{1, 2}
	v1.X = 1e9 // accessing strucut fields of v by dot notation
	fmt.Println(v1)

	p1 := &v1 // making pointer
	p1.Y = 69 // == (*p).Y = 77 but the language permits this short form
	fmt.Println(v1)

	v2 := Vertex{X: 10} // Y:0 is implicit
	p2 := &Vertex{2, 4}
	fmt.Println(v2, p2) // prefix & returns pointer to struct value
}
