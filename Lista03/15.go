package main

import (
	"fmt"
	"math"
)

func encontrardMinima(V []int) (int, int) {
	minDistancia := math.MaxInt64
	contagem := 0

	for i := 0; i < len(V)-1; i++ {
		for j := i + 1; j < len(V); j++ {
			d := intAbs(V[i] - V[j])
			contagem++
			if d < minDistancia {
				minDistancia = d
			}
		}
	}

	return minDistancia, contagem
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var N int
		fmt.Scan(&N)

		V := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&V[i])
		}

		minDistancia, contagem := encontrardMinima(V)

		fmt.Printf("%d %d\n", minDistancia, contagem)
	}
}
