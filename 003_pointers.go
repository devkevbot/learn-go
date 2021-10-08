package main

import "fmt"

func _003Pointers() {
	_int = 32
	_pointer := &_int
	fmt.Printf("Pointer value: %p\n", _pointer)
	fmt.Printf("Dereferenced value: %d\n", *_pointer)
}
