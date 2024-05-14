package main

import (
	"fmt"
	"strings"
)

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var CPF string
		fmt.Scan(&CPF)

		if validarCPF(CPF) {
			fmt.Println("CPF valido")
		} else {
			fmt.Println("CPF invalido")
		}
	}
}

func validarCPF(CPF string) bool {
	CPF = strings.ReplaceAll(CPF, " ", "")

	if len(CPF) != 11 {
		return false
	}

	for i := 0; i < 2; i++ {
		digito := calDigitoV(CPF[:9+i])
		digitoCPF := int(CPF[9+i] - '0')

		if digito != digitoCPF {
			return false
		}
	}

	return true
}

func calDigitoV(s string) int {
	var soma int
	for i := 0; i < len(s); i++ {
		soma += int(s[i]-'0') * (len(s) + 1 - i)
	}
	resto := soma % 11
	if resto < 2 {
		return 0
	}
	return 11 - resto
}
