package main

import "fmt"

func inverso (s string) string{
    n := []rune(s)
    for i, j :=0, len(n) -1; i < j; i, j = i+1, j-1 {
        n[i], n[j] = n[j], n[i]
    }
    return string(n)
}

func main() {
  var t int
  var num1, num2 string
  fmt.Scan(&t)
  
  for i := 0; i<t; i++{
    
      fmt.Scan(&num1)
      fmt.Scan(&num2)
    
      if inverso(num1)> inverso(num2){
          fmt.Print(inverso(num1),"\n")
      } else {
          fmt.Print(inverso(num2),"\n")
      }
    
   }
  
}
