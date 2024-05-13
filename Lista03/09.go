package main

import (
	"fmt"
	"math"
)

type Ponto struct {
	X float64
	Y float64
	Z float64
}

func distancia(p1, p2 Ponto) float64 {
	dX := p1.X - p2.X
	dY := p1.Y - p2.Y
	dZ := p1.Z - p2.Z

	return math.Sqrt(dX*dX + dY*dY + dZ*dZ)
}

func main() {
	var N int
	fmt.Scan(&N)

	Pontos := make([]Ponto, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&Pontos[i].X, &Pontos[i].Y, &Pontos[i].Z)
	}

	for i := 1; i < N; i++ {
		d := distancia(Pontos[i-1], Pontos[i])
		fmt.Printf("%.2f\n", d)
	}

}
