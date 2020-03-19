package main

import (
	"fmt"
	"strings"
)

func main() {
/*	var num float64;
	fmt.Print("请输入你需要转换的数字：")
	fmt.Scanf("%f",&num)
	//直接使用格式化输出%e或者%E。这时候就可以输出科学计数法
	fmt.Printf("科学计数法：%e\n",num)*/

	//字符串相加
	str1 := "hello"
	str2 := "world"
	str3 := str1 + str2
	fmt.Println(str3)
	//将数组切片（其实就是动态边长数组）
	var names = []string{"bill","steve","mark"}
	//在数组中插入一个新值
	names = append(names,"siryound")
	//在数组中的元素使用(字符串)连接起来。返回一个连接之后的大字符串。
	retStr := strings.Join(names,"&")
	fmt.Println(retStr)
	fmt.Printf("类型是%T",retStr)

}