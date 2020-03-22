package main

import (
	"fmt"
	"strconv"
)

/*
循环让用户输入人员信息，首先输入用户名，回车之后提示输入财富金额，如此循环往复
当用户输入over的时候，结束输入，并输入财富排行榜，包括幸福和财富
将这个游戏封装在game下，
 */
type person struct{
	name string
	money int
}//构造一个人的结构体
func main() {
	var ps []person //切片，是一个容器，就是结构体数组
	for {
		var name string
		var money string
		fmt.Print("请输入姓名:")
		fmt.Scanf("%s", &name)
		if name == "over"{
			break
		}
		fmt.Print("请输入财富：")
		fmt.Scanf("%s", &money)
		if money == "over" {
			break
		}
		ret,_ := strconv.Atoi(money)
		p := person{name, ret}
		ps = append(ps, p)
	}
	fmt.Println(ps)
	chooseSort(ps)
	for i,h := range(ps){
		fmt.Printf("位次：%d\n信息：%v\n",i+1,h)
	}
}

func chooseSort(ps []person) {
	for i:=0;i< len(ps)-1;i++  {
		ret := ps[i]
		index := i
		for j:=i+1;j<len(ps);j++{
			if(ps[j].money>ret.money){
				ret = ps[j]
				index = j
			}
		}
		ps[i],ps[index] = ps[index],ps[i]
	}
}
