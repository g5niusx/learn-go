package main

import "fmt"

func main() {
	// go中的基本数据类型和数组在函数中都是值传递，不会改变原来的值,在java中数组是引用传递会修改原来的值
	var day int8 = 8
	test(day)
	fmt.Println(day)
	// 对函数进行重命名
	var funcSum mySum = sum
	fmt.Printf("funcSum的类型 %T \n", funcSum)
	fmt.Printf("funcSum的运算结果 %v \n", funcSum(1, 3))
}

// 测试值传递
func test(day int8) {
	day--
}

func sum(a int8, b int8) int8 {
	return a + b
}

type mySum func(int8, int8) int8
