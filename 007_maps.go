package main

import (
	"fmt"
	"math"
)

type point struct {
	x float64
	y float64
}

func (p point) String() string {
	return fmt.Sprintf("(%.2f, %.2f)", p.x, p.y)
}

func (p point) isOnUnitCircle() bool {
	return math.Sqrt(p.x*p.x+p.y*p.y) == 1
}

func _007Maps() {
	/* Map literal */
	points := map[string]point{
		"origin":       {0, 0},
		"right point":  {1, 0},
		"top point":    {0, 1},
		"left point":   {-1, 0},
		"bottom point": {0, -1},
	}

	/* Maps can also be declared like so: make(map[string]int) */

	fmt.Printf("Interesting points:\n")
	for name, coord := range points {
		if coord.isOnUnitCircle() {
			fmt.Printf("%s is located on the unit circle\n", coord)
		} else {
			fmt.Printf("%s is located at: %s\n", name, coord)
		}
	}
}

func _007MutatingMaps() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
