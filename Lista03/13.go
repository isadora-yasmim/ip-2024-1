package main

import "fmt"

func main() {
	var N int
	maisFrequente := 0
	maiorFrequencia := 0

	fmt.Scan(&N)

	if N <= 1000 && N > 1 {
		V := make([]int, N)

		for i := 0; i < N; i++ {
			fmt.Scan(&V[i])
		}

		frequencias := make(map[int]int)

		for _, n := range V {
			frequencias[n]++
			if frequencias[n] > maiorFrequencia {
				maiorFrequencia = frequencias[n]
				maisFrequente = n
			}
		}
	}

	fmt.Printf("%d\n", maisFrequente)
	fmt.Printf("%d\n", maiorFrequencia)

}
