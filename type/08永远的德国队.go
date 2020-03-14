package main

import (
	"fmt"
	"time"
)

func main() {
	//无限死循环
/*	for {
	}*/
	var num int
	fmt.Println("请输入爱的次数：")
	fmt.Scan(&num)
	for i:=0 ; i < num; i++  {
		if i % 5 == 0 {
			fmt.Println("周五来了，最爱天台")
			continue
		}
		if i%2 == 0{
			fmt.Println("爱阿根廷")
		}else{
			fmt.Println("爱德国战车")
		}
		time.Sleep(time.Second)
	}
}