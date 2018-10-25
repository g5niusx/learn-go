package main

import (
	"fmt"
	"log"
	"time"
)

var (
	length = 50
)

func main() {
	// 数据channel
	dataChan := make(chan int, length)
	// 标志channel
	isCancel := make(chan bool, 1)
	// 向channel中写入数据
	go func(dataChan chan int) {
		for i := 0; i < length; i++ {
			dataChan <- i
			time.Sleep(time.Millisecond * 2)
			fmt.Printf("写入数据:%d\n", i)
		}
		//写完数据以后关闭channel
		close(dataChan)
	}(dataChan)
	// 从channel中读取数据
	go func(dataChan chan int, isCancel chan bool) {
		for {
			i, ok := <-dataChan
			if !ok {
				break
			}
			fmt.Printf("读取到数据:%d\n", i)
		}
		isCancel <- true
		close(isCancel)
	}(dataChan, isCancel)
	// 通过遍历一个标志channel来决定是否退出主线程
	for {
		_, ok := <-isCancel
		fmt.Printf("退出channel读取是否ok:%f\n", ok)
		if !ok {
			break
		}
	}

	// 申明一个channel为只写
	var writeOnlyChan chan<- int
	writeOnlyChan = make(chan int, 1)
	fmt.Printf("只写channel:%v\n", writeOnlyChan)

	// 申明一个channel为只读
	var readOnly <-chan int
	readOnly = make(chan int, 1)
	fmt.Printf("只读channel:%v\n", readOnly)

	intChan := make(chan int, length)
	for i := 0; i < length; i++ {
		intChan <- i
	}
	strChan := make(chan string, length)
	for i := 0; i < length; i++ {
		strChan <- fmt.Sprint(i)
	}
	// 在一个channel无法确定什么时候关闭的情况下。可以使用select来解决
test:
	for {
		select {
		case i := <-intChan:
			fmt.Printf("从int channel中获取到%v\n", i)
		case i := <-strChan:
			fmt.Printf("从string channel中获取到%v\n", i)
		default:
			fmt.Println("获取到默认值")
			//使用return会导致程序结束无法进行下面的代码
			//return
			break test
		}
	}
	// 在开启了协程以后，如果抛出异常，可以使用defer配合recover来处理，这样不会导致程序中断
	go panicFunc()
	// 休眠2秒，等待协程执行结束
	time.Sleep(time.Second * 2)
}

// 运行时候会抛出异常
func panicFunc() {
	defer func() {
		if recoverError := recover(); recoverError != nil {
			log.Fatalf("函数处理发生异常:%v", recoverError)
		}
	}()
	// 不make这个切片。来制造一个越界异常
	var sliceTest []string
	sliceTest[9] = "test"
}
