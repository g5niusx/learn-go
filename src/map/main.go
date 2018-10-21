package main

import "fmt"

// map是引用类型
func main() {
	var _map map[string]string
	_map = make(map[string]string)
	_map["1"] = "测试"
	_map["2"] = "test"
	fmt.Printf("%v,长度为:%v\n", _map, len(_map))

	cache := map[string]string{"code": "1"}
	fmt.Printf("%v\n", cache)
	// 删除key为code的元素
	delete(cache, "code")
	fmt.Printf("%v\n", cache)

	for key, value := range _map {
		fmt.Printf("%v %v\n", key, value)
	}

	// 切片map
	list := make([]map[string]int, 2)
	int1 := make(map[string]int)
	int1["age"] = 23
	int1["amount"] = 500
	list[0] = int1
	int2 := make(map[string]int)
	int2["age"] = 11
	int2["amount"] = 5000
	list[1] = int2
	list[1]["age"] = 40
	fmt.Printf("map类型的切片:%v,类型为:%T\n", list, list)

	// map排序
	sort := make(map[int]string)
	sort[2] = "测试1"
	sort[9] = "测试2"
	sort[4] = "测试3"
	sort[0] = "测试4"
	keys := make([]int, 0)
	for key := range sort {
		keys = append(keys, key)
	}
	// 对key的切片数组进行排序，然后排序输出
	bubbleSortNum(&keys)

	for _, value := range keys {
		fmt.Printf("%v", sort[value])
	}
	fmt.Println()
	personMap := make(map[int]Person)
	personMap[1] = Person{Name: "小王", Age: 23}
	personMap[2] = Person{Name: "小明", Age: 25}
	personMap[3] = Person{Name: "小李", Age: 30}
	fmt.Printf("结构体map:%v\n", personMap)
}

// 冒泡排序
func bubbleSortNum(keys *[]int) {
	for i := 0; i < len(*keys)-1; i++ {
		for j := 0; j < len(*keys)-1-i; j++ {
			if (*keys)[j] > (*keys)[j+1] {
				temp := (*keys)[j]
				(*keys)[j] = (*keys)[j+1]
				(*keys)[j+1] = temp
			}
		}
	}
}

type Person struct {
	Name string
	Age  int
}
