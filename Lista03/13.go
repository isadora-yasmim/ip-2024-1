package main
func main(){
	var N int
	fmt.Scan(&N)
	if N <= 1000 && N > 1 {
		V := make([]int, N)
		
		for i := 0; i < N; i++ {
			fmt.Scan(&V[i])
		}
	}

}