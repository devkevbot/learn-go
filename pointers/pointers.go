package main

import "fmt"

// Go has no pointer arithmetic

func main() {
	// p is a pointer to an int
	// var p *int

	// & generates a pointer to its operand
	// p = &i

	// * denotes the pointer's underlying value

	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

}
