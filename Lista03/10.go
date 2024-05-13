package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	notas := make([]int, N)
	var contagem, ultima, maior, indiceMaior int

	for i := 0; i < N; i++ {
		fmt.Scan(&notas[i])

		if i == N-1 {
			ultima = notas[i]
		}

		if notas[i] > maior {
			maior = notas[i]
			indiceMaior = i
		}
	}

	for _, nota := range notas {
		if nota == ultima {
			contagem++
		}
	}

	fmt.Printf("Nota %d, %d vezes\n", ultima, contagem)
	fmt.Printf("Nota %d, indice %d\n", maior, indiceMaior)
}
