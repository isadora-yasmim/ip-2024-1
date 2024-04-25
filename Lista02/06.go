package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Scan(&n)

	array := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}

	k := 1

	for i := 0; i < n-1; i++ {
		comp := 1 // comprimento da subsequência atual como 1
		for j := i + 1; j < n; j++ {
			if array[j] == array[i] {
				comp++
			} else {
				break
			}
		}

		if comp > k {
			k = comp
		}

	}

	fmt.Printf("A maior subsequência de elementos iguais possui %d elementos.", k)
}
