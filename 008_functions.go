package main

import "fmt"

func compute(x, y float64, fn func(float64, float64) float64) float64 {
	return fn(x, y)
}

func multiplier(x, y float64) float64 {
	return x * y
}

func _008FunctionValues() {
	fmt.Println("Computed value is:", compute(4, 2, multiplier))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func _008FunctionClosures() {
	pos, neg := adder(), adder()

	for i := 0; i < 5; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}
