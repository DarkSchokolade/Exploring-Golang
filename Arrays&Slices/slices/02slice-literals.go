// Do not worry  about the errors, because there needs to be only one main function
// This is for Practice only
package main

import "fmt"

func main() {
	// A slice literal is like an array literal without the length.
	r := []bool{true, false, false, true}
	r = append(r, true)
	fmt.Println(r)

	s := []struct {
		a int
		b bool
	}{
		{2, true},
		{4, true},
		{9, false},
	}
	fmt.Println(s)
}
