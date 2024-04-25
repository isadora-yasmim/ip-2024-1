package main

import "fmt"

func main() {
	var N int
	k := 1
	fmt.Scan(&N)

	if N < 2 {
		fmt.Println("Campeonato invalido!")
		return
	}

	for i := 1; i <= N-1; i++ {
		for j := i + 1; j <= N; j++ {
			fmt.Printf("Final %d: Time%d X Time%d\n", k, i, j)
			k++
		}
	}

}
