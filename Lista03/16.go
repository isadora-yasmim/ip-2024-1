package main

import "fmt"

func aulaCancelada(tempChegadas []int, K int) bool {
	cont := 0
	for _, tempChegada := range tempChegadas {
		if tempChegada <= 0 {
			cont++
		}
	}
	return cont < K
}

func M(tempChegadas []int) {
	for i := len(tempChegadas) - 1; i >= 0; i-- {
		if tempChegadas[i] <= 0 {
			fmt.Printf("%d\n", i+1) 
		}
	}
	fmt.Println()
}

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	var tempChegadas []int
	for i := 0; i < N; i++ {
		var tempChegada int
		fmt.Scan(&tempChegada)
		tempChegadas = append(tempChegadas, tempChegada)
	}

	if aulaCancelada(tempChegadas, K) {
		fmt.Println("SIM")
	} else {
		fmt.Println("NAO")
		M(tempChegadas)
	}
}
