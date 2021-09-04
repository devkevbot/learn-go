package main

import (
	"fmt"
	"runtime"
)

func main() {

	// switch statement will stop at first case that is true, no break statements needed
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}
