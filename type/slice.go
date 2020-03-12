package main

import "fmt"

func main() {
  var array = []int {0,1,2,3,4,5,6,7}
  s1 := array[0:4]
  s2 := array[:4]
  s3 := array[2:]
  fmt.Println(s1)
  fmt.Println(s2)
  fmt.Println(s3)

}
