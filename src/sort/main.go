package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// 排序
func main() {
	// 冒泡排序
	arrays := [...]int{23, 90, 1, 8, 79}
	slice := arrays[:]
	bubbleSort(&slice)

	names := [...]string{"biezhi", "os7blue", "legend", "小马", "小框", "g5niusx"}
	var inputName = ""
	fmt.Print("请输入需要查找的昵称:")
	fmt.Scanln(&inputName)

	// 顺序查找
	for index, value := range names {
		if strings.EqualFold(value, inputName) {
			fmt.Printf("在第[%d]位找到了该昵称[%s]\n", index+1, value)
			break
		} else if index == len(names)-1 {
			fmt.Printf("没有找到该昵称[%s]\n", inputName)
		}
	}

	// 二分查找(数组必须有序)
	bytes := randomArray(20)
	fmt.Printf("生成随机数组:%v\n", bytes)
	bubbleSort(&bytes)
	fmt.Printf("排序以后的数组:%v\n", bytes)
	value := bytes[rand.Intn(len(bytes))]
	fmt.Printf("随机查找的值为:%v\n", value)
	find := binaryFind(bytes, 0, len(bytes)-1, value)
	fmt.Printf("需要查找的值在第%v位", find+1)
}

// 二分查找
func binaryFind(arrays []int, left int, right int, value int) int {
	if right < left {
		return -1
	}
	middle := (left + right) / 2
	// 如果中间位的值等于查找的值说明已经找到
	if arrays[middle] == value {
		return middle
	} else if arrays[middle] < value {
		// 查到下标 middle+1 -> right 的范围
		return binaryFind(arrays, middle+1, right, value)
	} else {
		// 查找范围是下标 left -> middle-1
		return binaryFind(arrays, left, middle-1, value)
	}

}

// 生成一个随机切片
func randomArray(len int) []int {
	arrays := make([]int, len)
	for index := range arrays {
		rand.Seed(time.Now().UnixNano())
		arrays[index] = rand.Intn(50)
	}
	return arrays
}

// 冒泡排序
func bubbleSort(num *[]int) {
	for i := 0; i < len(*num)-1; i++ {
		for j := 0; j < len(*num)-1-i; j++ {
			// 如果i-1位大于i位则交换位置
			if (*num)[j] > (*num)[j+1] {
				temp := (*num)[j]
				(*num)[j] = (*num)[j+1]
				(*num)[j+1] = temp
			}
		}
		fmt.Printf("第%d次排序:%v\n", i, *num)
	}
}
