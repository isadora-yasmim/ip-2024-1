
package main
import "fmt"

func main() {
    s:=[]string{}
    a:=""
    for {
        fmt.Scan(&a)
        if a=="'"{
            break
        }
        s=append(s,a)
    }
    resultado := inverteFrase(s)
    fmt.Println(resultado)
}
func inverteFrase(t []string)(result string){
    for i := len(t) - 1; i >= 0; i-- {
        result += string(t[i])
    }
    return result
}
