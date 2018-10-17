package main

import "fmt"

// 切片是数组的引用
func main() {
	var arrays = [...]string{"g5niusx", "biezhi", "blue7os"}
	// 声明一个切片
	var slice = arrays[0:3]
	fmt.Printf("初始化的一个切片类型为:%T,切片里面的元素为为:%v,切片容器的大小为%v\n", slice, slice, cap(slice))
	fmt.Printf("切片的第一个元素内存地址为:%v,数组的第一个元素内存地址为:%v\n", &slice[0], &arrays[0])

	// 由于切片是数组的引用，测试修改数组是否会对切片产生影响
	arrays[0] = "test"
	// 输出的元素中有test
	fmt.Println(slice)

	//使用make来创建切片,第一个参数是切片的类型，第二个参数是切片的大小，第三个参数是切片的容量,切片的大小要小于等于切片的容量
	var ints []int = make([]int, 2, 10)
	fmt.Println(ints)
}
