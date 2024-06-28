
func SelectionSort(vetor []int) {
    n := len(vetor)
    for i := 0; i < n-1; i++ {
        menorIndice := i
        for j := i + 1; j < n; j++ {
            if vetor[j] < vetor[menorIndice] {
                menorIndice = j
            }
        }
        vetor[menorIndice], vetor[i]=vetor[i], vetor[menorIndice]
    }
}


