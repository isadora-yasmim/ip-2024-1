package main

import "fmt"

func main() {
	var N, primeiro int

	fmt.Scan(&N)
	fmt.Scan(&primeiro)
	fmt.Println(" ", primeiro)

	for i := 1; i < N; i++ {
		var num int
		fmt.Scan(&num)

		if num != primeiro {
			fmt.Println(" ", num)
			primeiro = num
		}
	}
}
