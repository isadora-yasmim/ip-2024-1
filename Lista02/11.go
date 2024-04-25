package main

import (
	"fmt"
	"math"
)

func main() {
	var n, e float64

	fmt.Scan(&n)
	fmt.Scan(&e)

	Rk := 1.0
	erro := math.Abs(n - Rk*Rk)

	iteraçao := 0
	for erro > e {

		Rk = (Rk + n/Rk) / 2
		erro = math.Abs(n - Rk*Rk)
		fmt.Printf("r: %.9f , erro: %.9f\n", Rk, erro)

		iteraçao++
	}

}
