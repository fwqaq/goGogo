package main

import "fmt"

func main() {
	//左移右移
	var num int = 1
	num = num << 1
	fmt.Println(num)
	num = num << 2
	fmt.Println(num)
	num = num >> 3
	fmt.Println(num)
	//异或
	fmt.Println(12^23)
	//或
	fmt.Println(34|45)
	//与运算
	fmt.Println(45&67)

}