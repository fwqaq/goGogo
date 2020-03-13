package main

import (
	"fmt"
	"time"
)

func doSomething(i int){
	fmt.Println("我是战狼小分队",i)
}

func main() {
	//小案例
	fmt.Println("总部要行动了！")//主协程

	for i:=1;i<101;i++{
		go doSomething(i)//在一条独立的携程中并发执行doSomething
	}
	fmt.Println("总司令晕过去了")
	time.Sleep(time.Second)//睡眠时间一秒钟
	fmt.Println("总司令醒过来了")
}
