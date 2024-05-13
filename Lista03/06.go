package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	v := make([]int, n)
	if n > 5000 {
		fmt.Println("Dígito inválido")
		return
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
	}

	soma := 0

	for i := 0; i < n; i++ {
		soma += v[i]
	}

	fmt.Print(soma)
}
