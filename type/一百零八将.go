package main

import "fmt"

/*声明布尔型*/
var isStupid bool = true//默认值是false

/*声明数值型*/
var age int//整型
var rmb float32//浮点型
var vector complex64 = 3+4i//复数

//声明字符串
var name string

/*声明复合型*/
var ages [10]int = [10]int{1,2,3,4,5,6,7,8,9,10}//10长度的整形数组
var height []int//不写长度，叫可边长的切片
var info map[string]float32 = map[string]float32{"资产":123.45,"负债":456}//映射，也就是字典键：值
var zh struct{
	name string
	age int
	rmb float32
}//结构体类型

var weight int = 300
var wieghtP *int = &weight//整形指针类型，存放的不是300，而是放置300的地址



/*声明函数*/
var getArea func()//无参无返回值

func main(){


	//fmt.Printf()是格式化的打印。%T是类型占位符，%v是值占位符。
	fmt.Printf("isStupid的类型是%T,值是%v\n",isStupid,isStupid)
	fmt.Printf("age的类型是%T,值是%v\n",age,age)
	fmt.Printf("rmb的类型是%T,值是%v\n",rmb,rmb)
	fmt.Printf("vector的类型是%T,值是%v\n",vector,vector)
	fmt.Printf("name的类型是%T,值是%v\n",name,name)
	fmt.Printf("ages的类型是%T,值是%v\n",ages,ages)
	fmt.Printf("height的类型是%T,值是%v\n",height,height)
	fmt.Printf("info的类型是%T,值是%v\n",info,info)
	fmt.Printf("zh的类型是%T,值是%v\n",zh,zh)
	fmt.Printf("wieghtP的类型是%T,值是%v\n",wieghtP,wieghtP)
	fmt.Printf("getArea的类型是%T,值是%v\n",getArea,getArea)

}