package main

import "fmt"

func getDj(score int)(string){
	if (score >= 90 && score <= 100)||(score < 10&& score > 0){
		return "优秀"
	}else if score >= 80{
		return "良"
	}else if score >= 70 {
		return "一般"
	}else {
		return "差"
	}
}

func main() {
	var score int
	fmt.Printf("骚年请输入你的分数：")
	fmt.Scan(&score)
	ret := getDj(score)
	fmt.Printf("您的等级是%s",ret)

}