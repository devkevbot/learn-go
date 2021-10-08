package main

import "fmt"

/* Basic built-in types */
var (
	_bool bool = false

	_string string = "Hello, world"

	_int   int   = -3
	_int8  int   = -127
	_int16 int16 = -1234
	_int32 int32 = -113_000
	_int64 int64 = -3_000_000

	_uint   uint   = 3
	_uint8  uint   = 127
	_byte   byte   = 128
	_uint16 uint16 = 1234
	_uint32 uint32 = 113_000
	_rune   rune   = 124_000
	_uint64 uint64 = 3_000_000

	_float32 float32 = 16_237_00.378
	_float64 float64 = 188_888_888_888.888

	_complex64  complex64  = 5 + 2i
	_complex128 complex128 = 128_312.1289 + 46_123.45i
)

/* Constants */
const (
	pi        = 3.14
	euler     = 2.781
	gravity   = 9.8
	world     = "世界"
	isRunning = false
)

func _001TypeInference() {
	inferredInt := 1
	fmt.Printf("Inferred type is %T, value is %d\n", inferredInt, inferredInt)
}

func _001TypeConversion() {
	i := 64
	f := float64(i)
	fmt.Printf("Original %T value is %d, converted %T value is %f\n", i, i, f, f)
}
