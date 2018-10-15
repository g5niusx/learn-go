package main

import (
	"fmt"
	"strings"
)

func main() {
	// go中的基本数据类型和数组在函数中都是值传递，不会改变原来的值,在java中数组是引用传递会修改原来的值
	var day int8 = 8
	test(day)
	fmt.Println(day)
	// 对函数进行重命名
	var funcSum mySum = sum
	fmt.Printf("funcSum的类型 %T \n", funcSum)
	fmt.Printf("funcSum的运算结果 %v \n", funcSum(1, 3))
	var sumAndSub2 sumSub = sumAndSub
	i, i2 := sumAndSub2(1, 3)
	fmt.Println(i, i2)
	// 可变参数
	fmt.Printf("可变参数的和为: %d \n", test2(1, 0, -9, -7, 8))
	// 交换位置
	var a = 11
	var b = 12
	swap(&a, &b)
	fmt.Println(a, b)
	// 匿名函数
	test := func(a int, b int) int {
		return a * b
	}(8, 9)
	fmt.Printf("匿名函数返回值 %v \n", test)
	// 测试闭包
	u := addUpper()
	fmt.Printf("闭包函数的值为: %v \n", u(1))
	fmt.Printf("闭包函数的值为: %v \n", u(1))
	// 闭包实践
	s := makeSuffix(".jpg")
	fmt.Printf("测试结果:%v \n", s("test"))
	fmt.Printf("测试结果:%v \n", s("自拍.jpg"))
}

// 测试值传递
func test(day int8) {
	day--
}

func sum(a int8, b int8) int8 {
	return a + b
}

func sumAndSub(a int, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return
}

func test2(num int, args ...int) int {
	count := num
	for i := 0; i < len(args); i++ {
		count += args[i]
	}
	return count
}

// 交换2个数的位置
func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

// 闭包
func addUpper() upper {
	// n申明在了函数中，对于返回的函数来说。n是不会改变的
	n := 10
	return func(i int) int {
		n += i
		return n
	}
}

// 传入一个后缀，如果传入的闭包参数带有后缀则直接返回，否则拼接后缀
func makeSuffix(suffix string) suffix {
	return func(s string) string {
		if strings.HasSuffix(s, suffix) {
			return s
		}
		s += suffix
		return s
	}
}

type mySum func(int8, int8) int8

type sumSub func(int, int) (int, int)

type upper func(int) int

type suffix func(string) string
