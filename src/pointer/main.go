package main

import "fmt"

func main() {
	var i = 10
	fmt.Printf("i 的内存地址是 %v \n", &i)
	// 声明了一个int类型的指针,里面真正存放的是i的内存地址
	var pointer *int = &i
	fmt.Printf("pointer存储的值为 %v,pointer所指向的值为 %v \n", pointer, *pointer)

	var num = 100
	numPointer := &num
	fmt.Printf("num的内存地址 %v \n", numPointer)
	// 通过指针修改num的值,通过*获取指针指向的值，然后再做运算或者做赋值
	*numPointer = 999
	fmt.Printf("num通过指针运算以后的值为 %v \n", num)
	// 指针指向不同的地址
	var a = 100
	var b = 300
	var ptr = &a
	*ptr = 999
	ptr = &b
	*ptr = 0
	fmt.Printf("a的值为 %v,b的值为 %v,指针的所指向的值为 %v \n", a, b, *ptr)
}
