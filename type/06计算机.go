package main

import "fmt"

func getResult(num1 float32,num2 float32,way string)(ret float32){
	if way=="+" {
		ret = num1+num2
	}else if way == "-" {
		ret = num1-num2
	}else if way=="*"{
		ret = num1*num2
	}else if way == "/" {
		ret = num1/num2
	}else if way == "%" {
		ret = float32(int(num1)%int(num2))
	}
	return
}
func getResultII(num1 float32,num2 float32,way string)(ret float32){
	switch way {
	case "+":
		ret = num1+num2
	case "-":
		ret = num1-num2
	case "*":
		ret = num1* num2
	case "/":
		ret = num1/num2
	case "%":
		ret = float32(int(num1)%int(num2))
	default:fmt.Printf("输入的格式不正确\n")
	}
	return
}
func main() {
	var way string
	fmt.Scanf("%s",&way)
	ret := getResult(5,3,way)
	fmt.Printf("结果ret=%v\n",ret)
	ret = getResultII(5,3,way)
	fmt.Printf("结果ret = %v\n",ret)

}

