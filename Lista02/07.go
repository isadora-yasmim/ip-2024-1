package main

import (
	"fmt"
)

func main() {
	var SImpar, SPar, n int
	totPar := 0
	totImpar := 0

	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		} else if n%2 == 0 {
			totPar++
			SPar = SPar + n

		} else {
			SImpar = SImpar + n
			totImpar++

		}
	}
	mi := SImpar / totImpar
	mp := SPar / totPar
	fmt.Printf("MEDIA PAR:%d\n", mp)
	fmt.Printf("MEDIA IMPAR:%d\n", mi)

}
