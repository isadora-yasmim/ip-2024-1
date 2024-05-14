package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	numeros := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&numeros[i])
	}

	sort.SliceStable(numeros, func(i, j int) bool {
		if numeros[i]%2 == numeros[j]%2 {
			if numeros[i]%2 == 0 {
				return numeros[i] < numeros[j]
			} else {
				return numeros[i] > numeros[j]
			}
		}
		return numeros[i]%2 < numeros[j]%2
	})

	for _, n := range numeros {
		fmt.Println(n)
	}
}
