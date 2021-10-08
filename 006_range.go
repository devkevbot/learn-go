package main

import "fmt"

func _006Range() {
	/* Range works with slices */
	powersOfTwo := []int{1, 2, 4, 8, 16, 32, 64, 128, 256}

	fmt.Println("Powers of two:")
	for index, value := range powersOfTwo {
		fmt.Printf("2**%d = %d\n", index, value)
	}

	/* Range also works with maps */
	_map := make(map[string]uint)
	_map["Apples"] = 2
	_map["Bananas"] = 4
	_map["Oranges"] = 1
	_map["Mangos"] = 0

	fmt.Println("Fruit count:")
	for key, value := range _map {
		fmt.Printf("%s: %d\n", key, value)
	}
}
