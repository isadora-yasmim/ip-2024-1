package main

import (
	"fmt"
)

func main() {
	array := []int{1, 2, 3, 4, 5}
	inverterArrayRec(array, 0, len(array)-1)
	fmt.Print(array)
}

func inverterArrayRec(arr []int, inicio int, fim int) {
	if inicio >= fim {
		return
	}
	arr[inicio], arr[fim] = arr[fim], arr[inicio]
	inverterArrayRec(arr, inicio+1, fim-1)
}
