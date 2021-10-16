package main

/* Returns the sum of all supplied integer arguments. */
func addIntegers(nums ...int) int {
	sum := 0
	for _, value := range nums {
		sum += value
	}
	return sum
}
