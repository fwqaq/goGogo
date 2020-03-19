package main

import "fmt"

//const Monday = 1
//const Tuesday = 2
//const Wed = 3
//const Thu = 4
//const Fri = 5
//const Sta = 6
//const Sun = 0
/*const (
	Sun = 0
	Monday = 1
	Tuesday = 2
	Wed = 3
	Thu = 4
	Fri = 5
	Sta = 6
)*/
//定义一组常量
const (
	Sun = iota
	Mon
	Tue
	Wed
	Thu
	Fri
	Sta
)
/*const (
	D = 1 << iota//变为1  相当于左移一位
	C = 1 << iota//下面的不写也可
	B = 1 << iota
	A = 1 << iota
	S = 1 << iota
)*/
const (
	D = 1 << iota//变为1  相当于左移一位
	C //下面的不写也可
	B
	A
	S
)
func main() {
	fmt.Println(Sun,Mon,Tue,Wed,Thu,Fri,Sta)
	fmt.Println(S,A,B,C,D)

}