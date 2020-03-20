package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
随机产生1000以内的数字
 */
func main() {
	randPow := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := randPow.Intn(1000)
	for  {
		fmt.Print("请输入你猜的数字:")
		var num int
		fmt.Scanf("%d",&num)
		if(num == i){
			fmt.Println("恭喜您，猜对了")
			break
		}else if num > i{
			fmt.Println("猜大了，是否继续，Y继续，N结束")
			str := "Y"
			fmt.Print("请输入是否继续游戏，Y继续，N结束：")
			fmt.Scanf("%s",&str)
			if str == "Y"{
				continue
			}else{
				fmt.Printf("游戏的正确答案是：%d",i)
			}
		}else{
			fmt.Println("猜小了，是否继续，Y继续，N结束")
			str := "Y"
			fmt.Print("请输入是否继续游戏，Y继续，N结束：")
			fmt.Scanf("%s",&str)
			if str == "Y"{
				continue
			}else{
				fmt.Printf("游戏的正确答案是：%d",i)
			}
		}
	}
}
