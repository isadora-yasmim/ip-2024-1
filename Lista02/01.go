package main

import (
	"fmt"
)

func main() {
	var M, F int
	var P, L, TF float64
	for{
		fmt.Scan(&M)
		if M == -1{
			break
		}
	var SP, SL float64
	for i := 0; i < 8; i++{
		fmt.Scan(&P)
		SP += P
	}
	for i := 0; i < 5; i++{
		fmt.Scan(&L)
		SL += L
	}
	fmt.Scan(&TF, &F)
	MP := SP / 8
	ML := SL / 5
	NF := 0.7 * MP + 0.15 * ML + 0.15 * TF

	situação := ""
	if NF >= 6 && F >= 0.75 * 128 {
		situação = "APROVADO"
	} else if NF >= 6 {
		situação = "REPROVADO POR FREQUENCIA"
	} else if F >= 0.75*128 {
		situação = "REPROVADO POR NOTA"
	} else {
		situação = "REPROVADO POR NOTA E POR FREQUENCIA"
	}
    fmt.Printf("Matricula: %d, Nota Final: %.2f, Situação Final: %s\n", M, NF, situação)
	}
}