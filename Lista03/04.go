package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n > 5000 {
		fmt.Println("Dígito inválido")
	}

	array := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])

	}

	fmt.Print(array)

}
