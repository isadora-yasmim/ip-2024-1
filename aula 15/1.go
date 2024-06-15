package main

import "fmt"

func main() {
	var x float64 = 3
	var n int = 2
	resultado := potenciaRec(x, n)
	fmt.Print(resultado)
}

func potenciaRec(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n < 0 {
		return 1 / potenciaRec(x, -n)
	}
	return x * potenciaRec(x, n-1)
}
