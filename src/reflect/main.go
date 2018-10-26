package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num int = 1
	reflectTest(&num)
	fmt.Printf("反射改变以后的值为:%v\n", num)
	reflectStruct()
}

func reflectTest(obj *int) {
	typeOf := reflect.TypeOf(obj)
	fmt.Printf("类型为:%v\n", typeOf)
	of := reflect.ValueOf(obj)
	fmt.Printf("值为:%v\n", of)
	kind := of.Kind()
	fmt.Printf("Kind为:%v\n", kind)
	// 是否允许做set操作
	set := of.CanSet()
	fmt.Printf("是否允许做set操作:%v\n", set)
	of.Elem().SetInt(40)
}

func reflectStruct() {
	person := Person{Name: "g5niusx", age: 18}
	typeOf := reflect.TypeOf(person)
	kind := typeOf.Kind()
	fmt.Printf("Kind为:%v\n", kind)
	fmt.Printf("结构体类型为:%v\n", typeOf)
	value := reflect.ValueOf(person)
	i := value.Interface()
	fmt.Printf("转化为interface:%v\n", i)
	fmt.Printf("获取到结构体的valueOf为:%v\n", value)
}

type Person struct {
	Name string
	age  uint
}
