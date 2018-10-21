package main

import "fmt"

// 结构体是值传递
func main() {
	var newCat *Cat = new(Cat)
	newCat.Age = 2
	fmt.Printf("使用new方式创建的对象:%v,修改年龄为:%v\n", newCat, newCat.Age)
	cat := Cat{Name: "花花", Age: 13}
	strings := make(map[string]string)
	strings["id"] = "0001"
	cat.info = strings
	fmt.Printf("初始化的内存地址为:%p\n", &cat)
	modify(cat)
	fmt.Printf("猫结构体的最终值为%v\n", cat.Name)
	ptr := &cat
	cat.pointer = ptr
	cat.friends = append(make([]Cat, 0), cat)
	fmt.Printf("最终的猫是:%v", cat)
}

func modify(cat Cat) {
	fmt.Printf("修改函数接收到的内存地址为:%p\n", &cat)
	cat.Name = "笑笑"
	cat.info["id"] = "0002"
	fmt.Printf("猫结构体在函数中被修改为%v\n", cat.Name)
}

type Cat struct {
	Age     int
	Name    string
	pointer *Cat
	friends []Cat
	info    map[string]string
}
