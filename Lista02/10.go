package main

import (
	"fmt"
)

func main() {
	var A int
	var B, C float64
	var S []float64
	var As []int
	var Bs, Cs []float64
	for {
		fmt.Scanf("%d %f %f\n", &A, &B, &C)
		if A == 0 {
			break
		}
		As = append(As, A)
		Bs = append(Bs, B)
		Cs = append(Cs, C)
	}
	for i := 0; i < len(As); i++ {
		S = append(S, Bs[i]*Cs[i])
	}

	for i := 0; i < len(S); i++ {
		fmt.Printf("%d %.2f \n", As[i], S[i])
	}
}
