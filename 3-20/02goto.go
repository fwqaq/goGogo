package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello")
	fmt.Println("golang")

	if time.Now().Hour() % 2== 1{
		fmt.Println("情绪稳定")
	}else{
		fmt.Println("糟了，情绪失常")
	}
	var i int
	for {
		if i > 10 {
			goto SHIF
		}
		time.Sleep(time.Millisecond*500)
		i++
	}
	fmt.Println("海景别墅")
	fmt.Println("满汉全席")
	SHIF:
		fmt.Println("此处免费领取房子")
}
