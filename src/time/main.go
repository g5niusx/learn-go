package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().Unix()
	now := time.Now()
	fmt.Printf("当前时间为:%v,类型为%T\n", now, now)
	fmt.Printf("当前年为:%v\n", now.Year())
	fmt.Printf("当前月为:%v\n", int(now.Month()))
	fmt.Printf("当前日为:%v\n", now.Day())
	fmt.Printf("当前时为:%v\n", now.Hour())
	fmt.Printf("当前分为:%v\n", now.Minute())
	fmt.Printf("当前秒为:%v\n", now.Second())
	fmt.Printf("当前纳秒为:%v\n", now.Nanosecond())

	// 格式化日期和时间
	format1 := now.Format("2006-01-02")
	fmt.Printf("%v格式化的时间为:%v \n", now, format1)

	//线程休眠
	time.Sleep(time.Second)

	//获取unix时间
	unix := now.Unix()
	nano := now.UnixNano()
	fmt.Printf("%v的unix秒为%v,纳秒为%v\n", now, unix, nano)

	//内建函数,new函数需要传入一个类型，返回一个初始化值的指针
	i := new(int)
	fmt.Printf("int类型分配的内存地址为%v,值为%v,类型为%T,指针所指向的值为%v\n", &i, i, i, *i)

	// 统计函数执行时间
	end := time.Now().Unix()
	fmt.Printf("当前函数的执行时间为%vs\n", end-start)
}
