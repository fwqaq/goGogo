package main

import "fmt"
const PI float32 = 3.14

func getArea(radius float32)(mianji float32){
	mianji = radius*radius*PI
	return
}
func getAreaII(){
	var radius float32
	var name string
	fmt.Print("请输入你的名字与脸半径:")
	fmt.Scanf("%s %f",&name,&radius)
	mianji := PI*radius*radius
	fmt.Printf("%s先生，你的脸有%.2f薄面\n",name,mianji)
}
func main(){

	var radius float32
	//让用户输脸的半径
	fmt.Print("骚年请输入你的脸的半径：")
	fmt.Scanf("%f",&radius)
	mianji := getArea(radius)
	fmt.Printf("你的脸有%.2f辣么大\n",mianji)
	getAreaII()
}