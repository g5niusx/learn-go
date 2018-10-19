package main

import (
	"fmt"
	"strconv"
	"strings"
)

const length = 10

// 二维数组
func main() {
	// 白棋
	white := "x"
	// 黑棋
	black := "y"
	arrays := [length][length]string{}
	_init(&arrays)
	printShape(arrays)
	fmt.Println("------ 棋盘初始化完毕 ------")
	fmt.Printf("[%s]代表白棋,[%s]代表黑棋,黑棋先手执行\n", white, black)
	fmt.Println("[1,1]表示将棋子下到左上角")
	count := 0
	for {
		if count%2 == 0 {
			fmt.Print("请黑棋落子>")
		} else {
			fmt.Print("请白棋落子>")
		}
		input := ""
		fmt.Scanln(&input)
		// 获取到数组
		x, y := getPoint(input)
		for !isValid(arrays, x, y) {
			fmt.Scanln(&input)
			x, y = getPoint(input)
		}

		// 判断是黑棋还是白棋
		if count%2 == 0 {
			arrays[x][y] = black
		} else {
			arrays[x][y] = white
		}
		count += 1
		// 打印棋盘
		printShape(arrays)
	}

}

// 判断棋子是否有效
func isValid(arrays [length][length]string, x int, y int) bool {
	if x > length || y > length {
		fmt.Println("数值不能大于5,请重新输入棋子位置")
		return false
	}
	if arrays[x][y] == "o" {
		return true
	}
	fmt.Println("不能够重复落子，请重新输入棋子位置")
	return false
}

// 初始化棋盘
func _init(arrays *[length][length]string) {
	for i := 0; i < len(arrays); i++ {
		for j := 0; j < len(arrays); j++ {
			arrays[i][j] = "o"
		}
	}
}

// 打印棋盘
func printShape(arrays [length][length]string) {
	for i := range arrays {
		for j := range arrays[i] {
			fmt.Print(arrays[i][j] + " ")
		}
		fmt.Println()
	}
}

// 获取下棋的坐标
func getPoint(input string) (x int, y int) {
	input = strings.TrimSpace(input)
	split := strings.Split(input, ",")
	x = str2int(split[0]) - 1
	y = str2int(split[1]) - 1
	return
}

// 字符串转数字
func str2int(str string) int {
	i, e := strconv.Atoi(str)
	if e != nil {
		fmt.Printf("字符串转数字异常:%v", e)
	}
	return i
}
