package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	V := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&V[i])
	}

	contagem := make(map[int]int)
	for _, num := range V {
		contagem[num]++
	}

	numElementosUnicos := 0
	for _, cont := range contagem {
		if cont == 1 {
			numElementosUnicos++
		}
	}

	
	fmt.Printf("%d\n", numElementosUnicos)
}