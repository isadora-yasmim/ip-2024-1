package main

import "fmt"

func main() {

	for {
		var tamanho int
		fmt.Scan(&tamanho)

		if tamanho == 0 {
			return
		}

		if tamanho <= 1 || tamanho > 10000 {
			fmt.Println("Valor inválido!")
			return
		}

		V := make([]int, tamanho)

		for i := 0; i < len(V); i++ {
			fmt.Scan(&V[i])

			if V[i] < 0 || V[i] > 10000 {
				fmt.Println("Valor inválido!")
				return
			}
		}

		M := 0

		for i := 0; i < len(V); i++ {
			if V[i] > V[M] {
				M = i
			}
		}

		contagem := 0

		for i := 0; i <= V[M]; i++ {
			for j := 0; j < len(V); j++ {
				if V[j] <= i {
					contagem++
				}
			}
			fmt.Printf("(%d) %d\n", i, contagem)
			contagem = 0
		}

		fmt.Printf("\n")
	}
}
