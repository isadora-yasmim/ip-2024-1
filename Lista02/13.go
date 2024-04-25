package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	tot := 0

	for i := 0; i < 64-1; i++ {
		// se o indice for par o quadro é escuro senão é branco
		if i%2 == 0 {
			tot += n * 2
		} else {
			tot += n
		}
	}

	fmt.Printf("%d\n", tot)
}
