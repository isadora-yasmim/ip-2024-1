package main

import (
	"fmt"
)

func main() {
	var n, i, k, s float64
	fmt.Scan(&n)
	fmt.Scan(&k)
	fmt.Scan(&i)
	fmt.Scan(&s)
	fmt.Println("")

	if n <= 9.0 && n >= 0.0 {
		fmt.Println("Tabuada de soma:\n")
		temp := i

		for a := 0.0; a < k; a++ {
			//tabuada de soma
			var R float64
			R = n + temp
			fmt.Printf("%.2f + %.2f = %.2f\n", n, temp, R)
			temp += s
		}
		fmt.Println("")
		fmt.Println("Tabuada de subtracao:\n")
		temp = i
		for a := 0.0; a < k; a++ {
			//tabuada de subtração
			var R float64
			R = n - temp
			fmt.Printf("%.2f - %.2f = %.2f\n", n, temp, R)
			temp += s

		}
		fmt.Println("")
		fmt.Println("Tabuada de multiplicacao:\n")
		temp = i
		for a := 0.0; a < k; a++ {
			//tabuada de multiplicação
			var R float64
			R = n * temp
			fmt.Printf("%.2f * %.2f = %.2f\n", n, temp, R)
			temp += s

		}
		fmt.Println("")
		fmt.Println("Tabuada de divisao:\n")
		temp = i
		for a := 0.0; a < k; a++ {
			//tabuada de divisão
			var R float64
			R = n / temp
			fmt.Printf("%.2f / %.2f = %.2f\n", n, temp, R)
			temp += s
		}

	}
}
