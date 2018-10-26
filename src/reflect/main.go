package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num = 1
	reflectTest(&num)
	fmt.Printf("反射改变以后的值为:%v\n", num)
	reflectStruct()

	person := Person{Name: "g5niusx", age: 18}
	reflectPerson := reflect.ValueOf(person)
	typePerson := reflect.TypeOf(person)
	// 获取到该实例有多少个属性
	numField := reflectPerson.NumField()
	fmt.Printf("%T有[%d]个字段\n", person, numField)
	for i := 0; i < numField; i++ {
		field := reflectPerson.Field(i)
		jsonTag := typePerson.Field(i).Tag.Get("json")
		fmt.Printf("%T第%d的字段值是:%v,jsonTag是%v\n", person, i, field, jsonTag)
	}
	// 获取该实例有多少个方法,私有的方法无法获取
	numMethod := reflectPerson.NumMethod()
	fmt.Printf("%T有[%d]个方法\n", person, numMethod)
	reflectPerson.Method(1).Call(nil)
	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf("golang")
	reflectPerson.Method(0).Call(params)
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
	Name string `json:"name"`
	age  uint
}

func (person Person) Say() {
	fmt.Printf("%v说:我是大撒比!!\n", person.Name)
}
func (person Person) Program(language string) {
	fmt.Printf("%v正在使用%v进行编程\n", person.Name, language)
}
