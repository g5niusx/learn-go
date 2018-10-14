package main

import "fmt"

func main() {
	fmt.Println("请输入星期几，来选择你的食谱;例如 1")
	var day int8
	fmt.Scanln(&day)
	switch day {
	case 1:
		fmt.Println("辣椒炒土豆丝")
	case 2:
		fmt.Println("麻辣小龙虾")
	case 3:
		fmt.Println("小酥肉")
	case 4:
		fmt.Println("炝炒娃娃菜")
	case 5:
		fmt.Println("火锅")
	default:
		fmt.Println("厨师休息!!!")
	}
}
