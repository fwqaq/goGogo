package main

import "fmt"
//九九乘法表
func jiujiu(){
	SHITsss:for i:=1;i<=9;i++{
		for j:=1;j<=i;j++{
			fmt.Printf("%d*%d=%d\t",i,j,i*j)
			if(j == 5&& i == 5){
				break SHITsss
			}
		}
		fmt.Println()
	}
}
//递归
func put10(num int){
	if(num < 0){
		return
	}
	fmt.Println(num)
	put10(num-1)
}
func main() {
	jiujiu()
	put10(10)
}
