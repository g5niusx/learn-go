package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var scores = [5]float64{}
	var test = [...]string{"g5niusx", "biezhi"}
	fmt.Printf("初始化数组 %v \n", test)
	var count = 0.0
	for i := 1; i < len(scores); i++ {
		fmt.Printf("请输入第[%d]的成绩 \n", i)
		score := 0.0
		// fmt.Scanln(&score)
		scores[i] = score
		count += score
	}
	fmt.Printf("总成绩是:%f,平均成绩是:%f \n", count, count/float64(len(scores)))

	// for-range遍历数组
	for index, value := range test {
		fmt.Printf("数组第%d的值是%v \n", index, value)
	}
	// 测试数组值传递
	test1(scores)
	fmt.Printf("值传递方法修改以后的值为:%v \n", scores[0])
	// 使用引用传递在方法里面修改值
	test2(&scores)
	fmt.Printf("引用传递方法修改以后的值为:%v \n", scores[0])
	fmt.Printf("构造的数组是%c \n", createArray())
	array := randomArray()
	fmt.Printf("生产随机数组:%v\n", array)
	fmt.Printf("反转以后的数组:%v\n", reverse(array))
}

func test1(scores [5]float64) {
	scores[0] = 90.2
}

func test2(scores *[5]float64) {
	scores[0] = 80
}

func createArray() [26]int {
	var arrays = [26]int{}
	for i := 0; i < len(arrays); i++ {
		arrays[i] = 65 + i
	}
	return arrays
}

// 创建一个随机int数组
func randomArray() [5]int {
	var arrays = [5]int{}
	rand.Seed(time.Now().Unix())
	for index, _ := range arrays {
		arrays[index] = rand.Intn(999)
	}
	return arrays
}

// 反转数组里面的元素
func reverse(arrays [5]int) [5]int {
	var copies = [5]int{}
	for index, _ := range arrays {
		copies[index] = arrays[len(copies)-index-1]
	}
	return copies
}
