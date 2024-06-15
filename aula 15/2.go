package main

import (
	"fmt"
)

func main() {
	array := []float64{1.2, 2.3, 3.4, 4.5, 5.6}
	resultado := somaRec(array, 0)
	fmt.Print(resultado)
}

func somaRec(arr []float64, index int) float64 {
	if index == len(arr) {
		return 0
	}
	return arr[index] + somaRec(arr, index+1)
}

