package main

import "fmt"

func main() {
	var N, M int
	fmt.Scanln(&N)

	V := make(map[int]bool)
	for i := 0; i < N; i++ {
		var num int
		fmt.Scan(&num)
		V[num] = true
	}

	fmt.Scanln(&M)

	for i := 0; i < M; i++ {
		var n int
		fmt.Scanln(&n)

		if V[n] {
			fmt.Println("ACHEI")
		} else {
			fmt.Println("NAO ACHEI")
		}
	}
}
