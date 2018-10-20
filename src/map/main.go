package main

import "fmt"

func main() {
	var _map map[string]string
	_map = make(map[string]string)
	_map["1"] = "测试"
	_map["2"] = "biezhi"
	fmt.Printf("%v,长度为:%v\n", _map, len(_map))

	cache := map[string]string{"code": "1"}
	fmt.Printf("%v\n", cache)
	// 删除key为code的元素
	delete(cache, "code")
	fmt.Printf("%v\n", cache)

	for key, value := range _map {
		fmt.Printf("%v %v\n", key, value)
	}

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
	fmt.Printf("map类型的切片:%v,类型为:%T", list, list)
}
