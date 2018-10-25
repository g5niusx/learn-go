package main

import "fmt"

var (
	resultMap = make(chan map[uint64]uint64, 0)
)

func main() {
	num := 200
	for i := 1; i < num; i++ {
		go Factorial(uint64(i))
	}
	result := uint64(0)
	for _map := range resultMap {
		fmt.Println(_map)
	}
	fmt.Printf("%d阶乘结果为%v\n", num, result)
}

// 求一个数的阶乘
func Factorial(num uint64) {
	result := uint64(1)
	for i := uint64(1); i <= num; i++ {
		result *= num
	}
	//向管道写入数据
	uint64s := make(map[uint64]uint64)
	uint64s[num] = result
	resultMap <- uint64s
}
