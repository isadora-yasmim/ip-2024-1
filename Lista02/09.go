package main

import "fmt"

func Primo(N int) bool {

	if N <= 1 {
		return false
	}

	if N <= 3 {
		return true
	}

	if N%2 == 0 || N%3 == 0 {
		return false
	}

	i := 5
	for i*i <= N {
		if N%i == 0 || N%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func main() {
	var N int
	_, err := fmt.Scan(&N)
	if err != nil || N <= 0 {
		fmt.Println("Numero invalido!")
		return
	}

	if Primo(N) {
		fmt.Println("PRIMO")
	} else {
		fmt.Println("NAO PRIMO")
	}
}
