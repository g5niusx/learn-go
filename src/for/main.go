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
}
