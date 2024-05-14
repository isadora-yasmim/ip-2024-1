package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	for {
		var N int
		fmt.Scan(&N)

		if N == 0 {
			break
		}

		V := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&V[i])
		}

		sort.Ints(V)

		fmt.Println(strings.Trim(fmt.Sprint(V), "[]"))
	}
}
