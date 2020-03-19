package main

import (
	"fmt"
	"unsafe"
)

func main(){
	var a int = 123
	b := 123
	var c = 123

	fmt.Printf("内存是%v\n",unsafe.Sizeof(a))
	fmt.Printf("内存是%v\n",unsafe.Sizeof(b))
	fmt.Printf("内存是%v\n",unsafe.Sizeof(c))

}
