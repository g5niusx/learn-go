package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

// 服务端
func main() {
	// 监听端口
	port := "23333"
	listen, e := net.Listen("tcp", ":"+port)
	if e != nil {
		log.Fatalf("连接端口[%s]异常:%v\n", port, e)
		return
	}
	defer listen.Close()
	log.Printf("开始监听端口[%s]....\n", port)
	for {
		conn, acceptError := listen.Accept()
		if acceptError != nil {
			log.Fatalf("监听端口[%s]异常:%v\n", port, e)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	log.Printf("客户端[%v]已经连接\n", conn.RemoteAddr().String())
	for {
		bytes := make([]byte, 1)
		n, err := conn.Read(bytes)
		if err == io.EOF || err == nil {
			// 因为申明来1024个字节。转化为字符串的时候需要使用返回的n来读取。确定有多少个字节
			fmt.Print(string(bytes[:n]))
		} else if err != nil {
			log.Fatalf("读取异常:%v\n", err)
			continue
		}
	}

}
