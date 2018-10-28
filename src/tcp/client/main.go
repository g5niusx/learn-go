package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

// 客户端
func main() {
	ip := "127.0.0.1:23333"
	conn, e := net.Dial("tcp", ip)
	if e != nil {
		log.Fatalf("连接[%s]异常:%v\n", ip, e)
		return
	}
	defer conn.Close()
	// 获取到标准输入
	reader := bufio.NewReader(os.Stdin)
	s, i := reader.ReadString('\n')
	if i != nil {
		log.Fatalf("获取输入异常:%v\n", i)
	}
	n, err := conn.Write([]byte(s))
	if err != nil {
		log.Printf("向[%s]写入数据异常%v\n", ip, err)
	} else {
		log.Printf("向[%s]写入[%d]数据\n", ip, n)
	}
}
