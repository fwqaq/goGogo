package main

import "fmt"

func main(){
  a := "hello"
  b := "world"

  c := a + b
  fmt.Println(len(a))
  fmt.Println(c)
  d := "hello,世界！"

  for i := 0; i < len(d); i++{
    fmt.Println(d[i])
  } 
  for i,v := range d{
    fmt.Println(i,v)
  }
  
}
