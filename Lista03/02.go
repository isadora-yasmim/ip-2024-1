package main

import "fmt"

func main() {
	var N, K int

	fmt.Scan(&N)

	if N <= 1 || N >= 1000 {
		fmt.Scan(&N)
	}

	V := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&V[i])
	}

	fmt.Scan(&K)

	Tot := 0
	for i := 0; i < N; i++ {
		if V[i] >= K {
			Tot++
		}
	}
	fmt.Println(Tot)

}
