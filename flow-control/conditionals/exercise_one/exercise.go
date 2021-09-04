package main

import "fmt"

// Sqrt blah
func Sqrt(x float64) float64 {
	z := 1.0
	lastZ := z

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println("last z is ", lastZ, "Z is ", z)

		if z == lastZ {
			break
		}

		lastZ = z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(100))
}
