package main

import "fmt"

func main(){

  //a := [3] int {1, 2, 3}//声明长度
  //a := []int{1, 2, 3}//不声明长度
  a := [3] int {1:2,2:4}//声明局部数据
  for i := 0;i < len(a); i++{
    fmt.Println(a[i])
  }
}
