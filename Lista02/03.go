package main

import "fmt"

func fatorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * fatorial(n-1)
}
func main() {
	var N int
	fmt.Println("Digite um número:")
	fmt.Scanln(&N)
	fmt.Printf("%d! = %d\n", N, fatorial(N))
}
