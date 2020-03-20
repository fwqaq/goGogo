package main

import "fmt"
//变长参数
func flexiable(a... interface{})  {
	//fmt.Println(a)
	for x,y := range a{
		fmt.Println(x,y)
	}
}
//有确定参数的不定长参数
func flexiable2(name string,age int,hobbys... string){
	fmt.Printf("我是%s，我%d大了，我的爱好是:",name,age)
	for index,value := range hobbys{
		fmt.Print(value)
		fmt.Println(index)
	}
}
func main() {
	flexiable(45,56,78,89)
	flexiable2("fw",23,"篮球","足球")
}
