package main

import (
	"fmt"
	//"strconv"
	"time"
)

func main() {
/*	//bool类型转换为string类型
	fmt.Println(strconv.FormatBool(true))
	//int转换为string
	i := strconv.FormatInt(789456,8)//第一个参数是数值，第二个是进制
	fmt.Println(i)
	//float转换为字符串，
	// 第一个参数:浮点型数据，
	// 第二个参数：format格式。可以使用b（二进制），e，E(e,E科学计数法的字符串)，f（浮点型字符串），g,G（g，G和精确度有关）
	// 第三个参数：精确度(f小数点后面保留几位)，如果是g，G那么表示有效数字的位数。
	// 第四个参数：结果用多少位表示，存储使用32位还是64位
	fmt.Println(strconv.FormatFloat(123.321,'g',5,32))
	//uint转换为字符串,第一次参数是数值，第二个是进制（）进制是输出的进制
	i = strconv.FormatUint(456489,10)
	fmt.Println(i)

	//字符串转换为基本类型
	//字符串转换为bool类型  返回第一个bool类型，
	// 第二个是error类型.如果成功值为nil。如果失败就是错误信息
	s1,err := strconv.ParseBool("true")
	if(err == nil){
		fmt.Println("转换成功，且值为",s1)
	}
	//字符串转换为int  第一个是字符串，
	// 第二个是进制，字符串表示的进制，
	// 第三个是存入数据内存大小
	s3,err := strconv.ParseInt("101",2,64)
	fmt.Println(s3,err)

	//字符串转换为float
	//第一个参数是浮点型，第二个参数是存储数据大小，也就是float32或者float64
	s2,err := strconv.ParseFloat("45621.321",64)
	if err == nil{
		fmt.Println(s2)
	}
	//字符串转换为uint
	s4,err := strconv.ParseUint("45675213",8,32)
	if err == nil{
		fmt.Println(s4)
	}

	//整形和字符串类型互换
	//整形转换为字符串
	fmt.Println(strconv.Itoa(45613))
	//字符串转换为整形  返回是整形和error
	fmt.Println(strconv.Atoi("456"))

	var ret interface{}//ret是任意类型
	ret = 1
	fmt.Println(ret)
	ret = "abc"
	fmt.Println(ret)*/

	//匿名函数
	go func(a,b int){
		fmt.Println(a+b)
	}(12,34)
	time.Sleep(500*time.Millisecond)
	fmt.Println("Game Over")

	var fuck = func(a,b int) int{
		return a+b
	}
	fmt.Println(fuck(45,89))
	fmt.Printf("类型是%T,数值是%v",fuck,fuck)
}
