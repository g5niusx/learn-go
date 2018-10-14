package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	// int8 占1个字节,范围是-2的7次方到2的7次方-1
	var int8Max int8 = (1 << 7) - 1
	var int8Min int8 = -1 << 7
	var i16Max int16 = (1 << 15) - 1
	var i16Min int16 = -1 << 15
	var i32Max int32 = (1 << 31) - 1
	var i32Min int32 = -1 << 31
	var i64Max int64 = (1 << 63) - 1
	var i64Min int64 = -1 << 63
	// unsafe.Sizeof可以返回变量所占的字节数
	fmt.Printf("类型为:%T,最小值为:%d,占用字节数为:%d,最大值为:%d,占用字节数为:%d\n", int8Min, int8Min, unsafe.Sizeof(int8Min), int8Max, unsafe.Sizeof(int8Max))
	fmt.Printf("类型为:%T,最小值为:%d,占用字节数为:%d,最大值为:%d,占用字节数为:%d\n", i16Min, i16Min, unsafe.Sizeof(i16Min), i16Max, unsafe.Sizeof(i16Max))
	fmt.Printf("类型为:%T,最小值为:%d,占用字节数为:%d,最大值为:%d,占用字节数为:%d\n", i32Min, i32Min, unsafe.Sizeof(i32Min), i32Max, unsafe.Sizeof(i32Max))
	fmt.Printf("类型为:%T,最小值为:%d,占用字节数为:%d,最大值为:%d,占用字节数为:%d\n", i64Min, i64Min, unsafe.Sizeof(i64Min), i64Max, unsafe.Sizeof(i64Max))
	var c1 byte = 'z'
	fmt.Printf("[%c]对应的码值为%d \n", c1, c1)
	var boolTrue = true
	fmt.Printf("[%v]占用字节数为:%d \n", boolTrue, unsafe.Sizeof(boolTrue))
	var name = "g5niusx"
	fmt.Println(name, &name)
	fmt.Printf("类型%T转化为类型%T \n", int8Min, string(int8Min))
	fmt.Println(fmt.Sprintf("int8类型的最小值是%d", int8Min))
	var boolStr = "True"
	b, _ := strconv.ParseBool(boolStr)
	fmt.Printf("boolStr转化后类型 %T,%v \n", b, b)
	var int8Str = "111"
	i, _error := strconv.ParseInt(int8Str, 2, 4)
	fmt.Printf("int8Str转化后类型 %T,%v 异常: %v \n", i, i, _error)
	float32Str := "123.099"
	f, _ := strconv.ParseFloat(float32Str, 64)
	fmt.Printf("float32Str转化后类型 %T,%v \n", f, f)
}
