package main

import "fmt"

//A struct is a collection of fields.

type Vertex struct {
	X int
	Y bool
}

func main() {
	fmt.Println(Vertex{1, false})

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
