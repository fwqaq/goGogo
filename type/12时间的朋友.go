package main

import (
	"fmt"
	"time"
)
func isRun(year int)bool{
	if(year%4==0&&year%100!=0) || (year%400==0){
		return true
	}
	return false
}

func main() {
	//获取当前时间的秒数，从1970 1 1
	elapseSeconds := int(time.Now().Unix())
	fmt.Println(elapseSeconds)
	elapseSeconds += 3600*8
	daystime := 3600*24
	yeartime := 3600*24*365
	var year int = 0
	var month int= 0
	var day int = 0
	xiaoshi := 0
	fenzhong := 0
	miaoshu := 0
	//算出年份
	for i:=1970;;i++ {
		if elapseSeconds < yeartime{
			year = i
			break
		}
		if isRun(i){
			elapseSeconds  = elapseSeconds - yeartime-daystime
		}else{
			elapseSeconds = elapseSeconds-yeartime
		}
	}
	//算出月份
	months := [13]int{0,31,28,31,30,31,30,31,31,30,31,30,31}
	if isRun(year){
		months[2] = 29
	}
	for i:=1;;i++{
		if elapseSeconds < months[i]*daystime {
			month = i
			break
		}
		elapseSeconds = elapseSeconds - months[i]*daystime
	}
	//算出日期
	for i:=1; ;i++  {
		if elapseSeconds < daystime{
			day = i
			break
		}
		elapseSeconds = elapseSeconds - daystime
	}
	//计算出小时  此时会有一个时差的问题，差距八小时的时差，在结果出会有八小时的时差
	xiaoshi = elapseSeconds / 3600
	elapseSeconds = elapseSeconds % 3600
	//计算出分钟数
	fenzhong = elapseSeconds / 60
	//计算出描述
	miaoshu = elapseSeconds % 60
	fmt.Printf("%v：%v：%v：%v：%v：%v",year,month,day,xiaoshi,fenzhong,miaoshu)
}
