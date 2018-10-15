package main

import "fmt"

func main() {
	// 打印矩形
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	// 打印一半矩形
	for i := 1; i <= 3; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	// 打印金字塔
	for i := 1; i <= 3; i++ {
		for k := 1; k <= 3-i; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	test(4)
	// 斐波那契数 0,1,1,2,3,5
	fmt.Println("斐波那契数:", test2(3))
	fmt.Println(test3(3))
	fmt.Println("桃子的总数:", feed(9))
}

func test(num int8) {
	if num > 2 {
		num--
		test(num)
	}
	fmt.Println(num)
}

// 斐波那契数
func test2(num int16) int16 {
	if num == 1 || num == 2 {
		return 1
	} else {
		return test2(num-1) + test2(num-2)
	}
}

func test3(num int8) int8 {
	if num == 1 {
		return 3
	}
	return 2*test3(num-1) + 1
}

// 🐒吃桃子问题,f(x) = (f(x+1)+1)*2
func feed(day int8) int8 {
	if day == 10 {
		return 1
	}
	return (feed(day+1) + 1) * 2
}
