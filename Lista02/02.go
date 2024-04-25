package main

import "fmt"

func main () {
	var A, B float64
     fmt.Println("Digite a população do país A= ")
	 fmt.Scanln(&A)
	 fmt.Println("Digite a população do país B= ")
	 fmt.Scanln(&B)
	 tempo := 0
	 for A < B {
		A = float64(A) * 1.03
		B = float64(B) * 1.015
		tempo++
	 }
	 fmt.Printf("O tempo gasto para que a população de A se igualasse a de B é de %d anos\n", tempo)

}