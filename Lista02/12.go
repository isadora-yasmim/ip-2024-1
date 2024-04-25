package main

import (
	"fmt"
)

func main() {
	var ValorIngresso, ValorInicial, ValorFinal float64

	fmt.Scan(&ValorIngresso)
	fmt.Scan(&ValorInicial)
	fmt.Scan(&ValorFinal)

	if ValorInicial >= ValorFinal {
		fmt.Println("INTERVALO INVALIDO")
		return
	}

	melhorValorFinal := 0.0
	melhorLucro := 0.0
	melhorNumIngressos := 0.0

	for v := ValorInicial; v <= ValorFinal; v += 1.0 {
		numIngressos := calcularNumIngressos(ValorIngresso, v)
		lucro := calcularLucro(numIngressos, ValorIngresso, v)

		if lucro > melhorLucro {
			melhorLucro = lucro
			melhorValorFinal = v
			melhorNumIngressos = numIngressos
		}

		fmt.Printf("V: %.2f, N: %.f, L: %.2f\n", v, numIngressos, lucro)

	}

	fmt.Printf("Melhor valor final: %.2f\n", melhorValorFinal)
	fmt.Printf("Lucro: %.2f\n", melhorLucro)
	fmt.Printf("Numero de ingressos: %.f\n", melhorNumIngressos)
}

func calcularNumIngressos(ValorIngresso, v float64) float64 {
	var numIngressos float64

	if v < ValorIngresso {
		numIngressos = 120.0 + ((ValorIngresso - v) / 0.5 * 25)
	} else if v > ValorIngresso {
		numIngressos = 120.0 - ((v - ValorIngresso) / 0.5 * 30)
	} else {
		numIngressos = 120.0
	}

	return numIngressos
}

func calcularLucro(numIngressos float64, ValorIngresso, v float64) float64 {
	gastosfixos := 200.0 + 0.05*float64(numIngressos)
	receita := float64(numIngressos) * v

	return receita - gastosfixos
}
