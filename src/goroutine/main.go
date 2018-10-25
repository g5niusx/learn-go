package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

var (
	resultMap = make(map[uint64]uint64)
	// 申明一个互斥锁
	lock sync.Mutex
)

func main() {
	// 当前系统有多少个逻辑cpu
	cpu := runtime.NumCPU()
	// 设置多少个cpu来执行任务
	runtime.GOMAXPROCS(cpu)
	fmt.Printf("当前系统有%d个CPU\n", cpu)
	// go关键字。开启来一个协程,不会阻塞
	go hello(1)
	defer fmt.Println("defer 调用。。。。")
	for i := 0; i < 1; i++ {
		fmt.Println("主线程:hello world")
		time.Sleep(time.Second)
	}
	num := 200
	for i := 1; i < num; i++ {
		go Factorial(uint64(i))
	}
	result := uint64(1)
	lock.Lock()
	for _, value := range resultMap {
		result += value
	}
	lock.Unlock()
	fmt.Printf("%d阶乘结果为%v\n", num, result)
	// 主线程退出以后，协程也会退出,直接调用exit退出程序。并不会触发defer关键字
	os.Exit(0)
}

func hello(num int) {
	for i := 0; i < num; i++ {
		fmt.Println("函数:hello world")
		time.Sleep(time.Second)
	}
}

// 求一个数的阶乘
func Factorial(num uint64) {
	result := uint64(1)
	for i := uint64(1); i <= num; i++ {
		result *= num
	}
	// 对于写map操作加锁
	lock.Lock()
	resultMap[num] = result
	lock.Unlock()
}
