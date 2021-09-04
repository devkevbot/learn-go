package main

import "fmt"

// a struct is a collection of fields
type Vertex struct {
	X int
	Y int
}

func main() {

	// print out struct fields at once
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}

	// print out the X field
	v.X = 4
	fmt.Println(v.X)

	// pointers to structs are allowed
	p := &v

	// we can also write (*p).X, but explicit dereferencing is not needed
	p.X = 1e9
	fmt.Println(v)

	// struct literals are allowed
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		g  = &Vertex{1, 2} // has type *Vertex
	)

	fmt.Println(v1, g, v2, v3)
}
