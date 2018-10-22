package main

import "fmt"

// 类型断言
func main() {
	var test interface{}
	a := Point{x: 10, y: 20}
	test = a
	var b Point
	b, ok := test.(Point)
	if ok {
		fmt.Println("类型转化成功")
		fmt.Printf("类型断言:%v\n", b)
	}
}

type Point struct {
	x float64
	y float64
}
