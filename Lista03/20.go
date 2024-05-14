package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)

	var maiorX, maiorY, maiorZ float64

	var x1, y1, z1 float64
	fmt.Scan(&x1, &y1, &z1)

	for i := 1; i < N; i++ {
		var x2, y2, z2 float64
		fmt.Scan(&x2, &y2, &z2)

		VX := math.Abs(x2 - x1)
		VY := math.Abs(y2 - y1)
		VZ := math.Abs(z2 - z1)

		if VX > maiorX {
			maiorX = VX
		}
		if VY > maiorY {
			maiorY = VY
		}
		if VZ > maiorZ {
			maiorZ = VZ
		}

		x1, y1, z1 = x2, y2, z2
	}

	fmt.Printf("%.2f\n%.2f\n%.2f\n", maiorX, maiorY, maiorZ)
}
