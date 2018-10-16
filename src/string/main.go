package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "g5niusx我"
	//获取字符串长度,获取的是utf-8的unicode,中文会占用3个字节长度
	fmt.Printf("字符串长度 %v \n", len(str))
	//获取中文的字符串长度
	r := []rune(str)
	fmt.Printf("切片字符串长度 %v \n", len(r))
	// 字符串转整数
	i, e := strconv.Atoi("test")
	if e != nil {
		fmt.Printf("字符串转整数异常 %v \n", e)
	} else {
		fmt.Printf("字符串转整数 %v \n", i)
	}

	//字符串转byte
	_byte := []byte(str)
	fmt.Printf("字符串转byte: %v \n", _byte)

	//byte转字符串
	fmt.Printf("byte转字符串: %s \n", string(_byte))

	// 10进制转制定的进制
	formatInt := strconv.FormatInt(900, 2)
	fmt.Printf("900转2进制是 %v \n", formatInt)
	fmt.Printf("900转16进制是 %v \n", strconv.FormatInt(900, 16))

	//字符串里面是否包含指定的字符串
	source := "ceshiceshi"
	target := "ce"
	fmt.Printf("%s是否包含%s:%v \n", source, target, strings.Contains(source, target))

	// 字符串里面包含多少个指定的字符串
	fmt.Printf("%s包含%d个%s \n", source, strings.Count(source, target), target)

	// 获取字符串在指定字符串第一次出现的位置
	index := strings.Index(str, string(r[1]))
	fmt.Printf("字符[我]出现在%v的位置是 %d \n", str, index)

	// 将指定的字符串替换
	replace := strings.Replace(str, "我", "|biezhi", -1)
	fmt.Printf("%v替换后的结果是%v \n", str, replace)
	str = "123,456,890"

	// 将字符串按照指定字符分割为字符串数组
	split := strings.Split(str, ",")
	fmt.Printf("%v分割以后的数组是: %v \n", str, split)

	// 将字符串大写
	str = "upperCase"
	upper := strings.ToUpper(str)
	fmt.Printf("%v转换为大写以后为%v \n", str, upper)

	// 去掉字符串首尾空格
	str = " g5niusx\n"
	space := strings.TrimSpace(str)
	fmt.Printf("%v去掉首尾空格以后为%v \n", str, space)
}
