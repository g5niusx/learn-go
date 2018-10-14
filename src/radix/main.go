package main

import "fmt"

// 进制测试
func main() {
	var i = 2
	// 二进制
	fmt.Printf("%v的二进制是%b \n", i, i)
	// 八进制go中以0开头的整数表示为8进制
	var j = 01101
	fmt.Println(j)
	// 十六进制
	var z = 0xF1109
	fmt.Println(z)
	trans(30, 0)
}

func trans(i int, radix int) {
	str := string(i)
	fmt.Println(len(str))
}
