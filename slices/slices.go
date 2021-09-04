package main

import "fmt"

func main() {

	// Slices are dynamically-sized
	// []T is a slice with elements of type T
	// a[low:high] is half-open; includes low up to high-1

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	// Slices are like references to arrays
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// Changing the elements of a slice modifies the corresponding elements of its underlying array
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
