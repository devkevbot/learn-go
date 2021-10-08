package main

import "fmt"

func _005Array() {
	/* All set to 8 */
	var key [16]uint8
	fmt.Printf("Array of 16 integers: %v\n", key)

	names := [3]string{"Bob", "Tom", "Mary"}
	fmt.Printf("Array of 3 strings: %v\n", names)
}

func _005Slice() {
	names := [3]string{"Bob", "Tom", "Mary"}
	/* Range on slices is in the form [start, end) */
	firstTwoNames := names[0:2]
	fmt.Printf("First and second names: %v\n", firstTwoNames)

	/* Changing the slice value alters the array value that is being referenced! */
	firstTwoNames[0] = "Terry"
	fmt.Printf("First and second names after mutation: %v\n", firstTwoNames)

	/* Slice literals */
	numbers := []int{1, 2, 3}
	fmt.Printf("Slice literal of ints: %v\n", numbers)

	/* Slices have a length AND a capacity
	Length = # of elements in the slice
	Capacity = # of elements in the underlying array
	*/

	_slice := []int{3, 7, 9, 10, 15}
	printSlice(_slice)

	_slice = _slice[:3]
	printSlice(_slice)

	/* Create a slice using make() */
	_makeSlice := make([]int, 10)
	printSlice(_makeSlice)

	/* Slices of slices */
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	fmt.Printf("Board is of type %T\n", board)

	/* Appending to a slice */
	trophyCase := []string{"1st", "1st", "2nd"}
	trophyCase = append(trophyCase, "4th", "2nd", "1st")
	fmt.Printf("Appended to a the trophy case: %v\n", trophyCase)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
