package main

import (
	"fmt"
)

// The var statement declares a list of variables; as in function argument lists, the type is last.
var c, python, java bool

var a, b int = 1, 2

func main() {
	var i int

	/*
		A var declaration can include initializers, one per variable.
		If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
	*/
	var cat, dog = "meow!", "bark!"

	/*
		Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
		Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
	*/
	k := 3

	fmt.Println(i, c, python, java)
	fmt.Println(cat, dog)
	fmt.Println(k)
}
