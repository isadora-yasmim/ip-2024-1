package main

import "fmt"

func main() {
	for {
		var N int
		fmt.Scan(&N)

		if N == 0 {
			break
		}

		v := make([]int, N)

		for i := 0; i < N; i++ {
			fmt.Scan(&v[i])
		}

		vMax := v[0]
		iMax := 0

		for i := 1; i < N; i++ {
			if v[i] > vMax {
				vMax = v[i]
				iMax = i
			}
		}

		fmt.Println(iMax, vMax, " ")
	}
}
