package main

import (
	"encoding/json"
	"fmt"
)

// 结构体是值传递
func main() {
	var newCat = new(Cat)
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
	fmt.Printf("最终的猫是:%v\n", cat)

	a := A{No: 1}
	b := B{No: 2}
	a = A(b)
	a.test("test")
	fmt.Printf("b对象强转为a:%v\n", a)

	// 结构体序列化
	bytes, _ := json.Marshal(cat)
	jsonString := string(bytes)
	fmt.Printf("cat序列化以后的字符串为:%v\n", jsonString)

	person := Person{Name: "os7blue"}
	person.goodMan()
}

func modify(cat Cat) {
	fmt.Printf("修改函数接收到的内存地址为:%p\n", &cat)
	cat.Name = "笑笑"
	cat.info["id"] = "0002"
	fmt.Printf("猫结构体在函数中被修改为%v\n", cat.Name)
}

type Cat struct {
	Age     int    `json:"age"`
	Name    string `json:"name"`
	pointer *Cat
	friends []Cat
	info    map[string]string
}

type A struct {
	No int
}

func (a A) test(name string) {
	fmt.Printf("A对象的编号是:%v\n", a.No)
}

type B struct {
	No int
}

type Person struct {
	Name string
}

func (person Person) goodMan() {
	fmt.Printf("%s,你是一个好人", person.Name)
}
