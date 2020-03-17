package main

import (
	"fmt"
	"math/rand"
)

func main() {
	 computers := [3]string{"石头","剪刀","布"}
	fmt.Println("请出拳：石头，剪刀，布")
	//同时定义两个变量
	var userFist,computer string
	cf := rand.Intn(3)

	computer = computers[cf]

	fmt.Scanf("%s",&userFist)
	switch userFist {
	case "石头":
		if computer == "石头" {
			fmt.Println("平局")
		}else if computer == "剪刀" {
			 fmt.Println("您胜了")
		}else if computer == "布" {
			fmt.Println("电脑胜了，你输了")
		}
		fmt.Printf("您出的是%s,电脑出的是%s",userFist,computer)
	case "布":
		if computer == "布" {
			fmt.Println("平局")
		}else if computer == "石头"{
			fmt.Println("你胜了")
		}else {
			fmt.Println("你输了，电脑胜了")
		}
		fmt.Printf("您出的是%s,电脑出的是%s",userFist,computer)

	case "剪刀":
		if computer == "剪刀"{
			fmt.Println("平局")
		}else if computer == "石头"{
			fmt.Println("电脑胜了，你输了")
		}else{
			fmt.Println("你胜了")
		}
		fmt.Printf("您出的是%s,电脑出的是%s",userFist,computer)

	default:
		fmt.Println("非法输入")
	}


}
