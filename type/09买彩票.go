package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for {
		fmt.Println("抽奖开始了，人生巅峰即将开始，心里有点小激动....")
		fmt.Println("请输入你要的购买的金额和队伍空格隔开")
		var money int
		var duiwu string
		fmt.Scanf("%d+%s", &money, &duiwu)
		//获得一个最新的种子
		rand.New(rand.NewSource(time.Now().Unix()))
		//创建一个随机数
		value := rand.Intn(101)
		if value > 90 {
			fmt.Println("你买的%s队伍已经胜利，您获胜了，共获得金额%v", duiwu, money*100)
			fmt.Println("靠海别墅")
		} else {
			fmt.Println("你没有买中，下此继续努力")
			fmt.Println("大厦天台")
		}
	}

}