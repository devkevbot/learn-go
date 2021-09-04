package main

import "fmt"

func main() {

	// the FOR loop is the only looping construct in GO
	/*
		sum := 0
		for i := 0; i < 10; i++ {
			sum += i
			fmt.Println(sum)
		}
	*/

	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// infinite loop
	for {

	}
}
