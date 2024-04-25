//fiz a questão 13 dessa outra forma também, que não percorre o tabuleiro,po´rem tambpem da pra calcular o tot de grãos
package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scan(&n)

	tot := (n * 31) + (2 * n * 32)
	//(n * 31)para os quadros brancos + (2 * n * 32) para os quadros escuros
	fmt.Printf("%d", tot)
}