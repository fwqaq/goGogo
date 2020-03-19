package main

import "fmt"

func main() {
	numsArray := getNumbers(456)
	fmt.Println(numsArray)
}

func getNumbers(i int) (nums [3]int) {
	//123提取出来
	for j:=2; j >= 0; j--  {
		nums[j] = i%10
		i/=10
	}
	return
}
