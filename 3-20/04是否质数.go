package main

import (
	"fmt"
	"math"
)

func isPrimeNumber(num int)(bool){
	if num == 1{
		return false
	}
	i := math.Sqrt(float64(num))
	for j:=2;j<=int(i);j++{
		if num % j == 0{
			return false
		}
	}
	return true
}

func main() {
	var num int = 0
	for i:=2;i<1000;i++  {
		if isPrimeNumber(i){
			fmt.Println(i)
			num++
		}
	}
	fmt.Println("总共有%v素数",num)
}
