package main

import (
	"fmt"
	"strconv"
)

func main() {
	numero := 10
	binario := decimalParaBinario(numero)
	fmt.Print(binario)
}

func decimalParaBinario(n int) string {
	if n == 0 {
		return "0"
	}
	if n == 1 {
		return "1"
	}
	return decimalParaBinario(n/2) + strconv.Itoa(n%2)
}
