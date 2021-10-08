package main

import "fmt"

func _010TypeAssertions() {
	var i interface{} = "Hello"

	s, ok := i.(string)
	fmt.Println("OK?", ok, "->", s)

	f, ok := i.(float64)
	fmt.Println("OK?", ok, "->", f)
}

func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("v squared is %d\n", v*v)
	case string:
		fmt.Printf("string of length %d\n", len(v))
	default:
		fmt.Println("Neither an int nor a string")
	}
}

func _010TypeSwitches() {
	checkType(10)
	checkType("Hello, world")
	checkType(false)
}
