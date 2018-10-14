package main

import "fmt"

func main() {
	var name string
	var age int8
	var score float32
	fmt.Println("请输入姓名...")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄...")
	fmt.Scanln(&age)
	fmt.Println("请输入分数...")
	fmt.Scanln(&score)
	fmt.Println("--------------------")
	fmt.Printf("姓名是:%v,年龄是%v,分数是%v \n", name, age, score)
	fmt.Println("--------------------")

	fmt.Println("使用更简单的方法搞定😄")
	fmt.Println("请输入对应的姓名，年龄，分数；使用空格来分割，例如：g5niusx 23 60")
	fmt.Scanf("%s %d %f", &name, &age, &score)
	fmt.Println("--------------------")
	fmt.Printf("姓名是:%v,年龄是%v,分数是%v \n", name, age, score)
	fmt.Println("--------------------")
}
