package main
import "fmt"

func main() {
    a := [5]int{1, 2, 3, 4, 5}
    b := [5]int{6, 7, 8, 8, 9}
    
    aSlice := a[:]
    bSlice := b[:]
    
    resultado:=trocaOpostosSeMenor(aSlice, bSlice)
    fmt.Println(aSlice, bSlice, resultado)
    }
    
func trocaOpostosSeMenor(x, y []int)bool{
    trocou:= false
    for i:= 0;i<len(x);i++{
        for j:=len(y)-1;j>i;j--{
            if x[i]<y[j]{
                x[i], y[j]=y[j],x[i]
                trocou= true
            }
        }
    }
    return trocou
}
