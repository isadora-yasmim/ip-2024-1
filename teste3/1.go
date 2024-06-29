package main
import "fmt"

func main() {
    s:=""
    fmt.Scan(&s)
    resultado := inverteFrase(s)
    fmt.Println(resultado)
}
func inverteFrase(str string)(result string){
    for _, t:= range str{
        result = string(t)+ result
    }
    return result
}
