package main

import "fmt"

func main() {
	// When slicing, you may omit the high or low bounds to use their defaults instead
	// Default is 0 for the low bound and the length of the slice for the high bound

	// var a [10]int
	// is equivalent to
	// a[0:10]
	// a[:10]
	// a[0:]
	// a[:]

	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
