package main

import "fmt"

func main() {
	var a int
	j := a            // j is an int
	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128
	v := 42           // change me!
	fmt.Printf("v is of type %T\n", v)
}
