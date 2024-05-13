package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		var N int32
		fmt.Scan(&N)

		binary := strconv.FormatInt(int64(N), 2)
		
		fmt.Println(binary, "\n")
	}
}
