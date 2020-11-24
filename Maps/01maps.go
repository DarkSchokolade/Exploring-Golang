package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

//	A map maps keys to values.
//	A nil map has no keys, nor can keys be added.
//	Use make to return a map of given type.
var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	var mapLiterals = map[string]Vertex{
		"Eiffiel tower": Vertex{76.86575, -32.95844},
		"Dominos":       {24.55552, 45.12356}, //the Vertex part is optional
	}
	fmt.Println(mapLiterals)
}
