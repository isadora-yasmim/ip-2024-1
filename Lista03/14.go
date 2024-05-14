package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int

	fmt.Scan(&N)

	if N <= 1000000 && N > 0 {
		V := make([]int, N)

		for i := 0; i < N; i++ {
			fmt.Scan(&V[i])
		}

		sort.Ints(V)

		mediana := 0.0

		if N%2 == 0 {
			mediana = float64(V[N/2-1]+V[N/2]) / 2.0
		} else {
			mediana = float64(V[N/2])
		}

		fmt.Printf("%.2f\n", mediana)

	}
}
