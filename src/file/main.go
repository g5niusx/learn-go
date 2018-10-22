package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const (
	PATH = "/Users/g5niusx/Workspace/learn-go/src/file/test.txt"
)

func main() {
	path := "/Users/g5niusx/Workspace/learn-go/src/file/main.go"
	file, e := os.Open(path)
	if e != nil {
		fmt.Printf("打开文件异常%v", e)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var content string
	for {
		s, i := reader.ReadString('\n')
		// 表示结束了
		if i == io.EOF {
			break
		}
		content = content + s
		if i != nil {
			fmt.Printf("读取文件异常:%v", i)
		}

	}
	fmt.Printf("读取到内容如下\n%v", content)
	fmt.Println("-------------使用utils一次性读取文件内容---------------------")
	bytes, _ := ioutil.ReadFile(path)
	fmt.Printf("%v", string(bytes))

	openFile, i := os.OpenFile(path, os.O_RDWR, os.ModeDir)
	if i != nil {
		fmt.Printf("打开%s异常%v", path, i)
	}
	defer openFile.Close()
	newReader := bufio.NewReader(openFile)
	content = ""
	for {
		s, i := newReader.ReadString('\n')
		// 表示结束了
		if i == io.EOF {
			break
		}
		content = content + s
		if i != nil {
			fmt.Printf("读取文件异常:%v", i)
		}
	}
	fmt.Printf("文件内容\n%v", content)
	createWrite(content)
	override("测试覆盖文件\r\n")
	_append("123456")
	// 删除文件
	remove := os.Remove(PATH)
	if remove != nil {
		fmt.Printf("删除[%s]文件异常:%v", PATH, remove)
	}
}

// 创建文件并且写入内容
func createWrite(content string) {
	// 默认的创建方式，任何人都可读写，不可执
	file, e := os.OpenFile(PATH, os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	if e != nil {
		fmt.Printf("创建[%s]异常:%v", PATH, e)
		return
	}
	writer := bufio.NewWriter(file)
	// 写入文件
	writer.WriteString(content)
	// 刷新内容
	writer.Flush()

}

// 覆盖文件内容
func override(content string) {
	// 打开的时候清空文件
	file, e := os.OpenFile(PATH, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	defer file.Close()
	if e != nil {
		fmt.Printf("创建[%s]异常:%v", PATH, e)
		return
	}
	writer := bufio.NewWriter(file)
	// 写入文件
	writer.WriteString(content)
	// 刷新内容
	writer.Flush()
}

// 追加内容
func _append(content string) {
	// 打开的时候清空文件
	file, e := os.OpenFile(PATH, os.O_APPEND|os.O_WRONLY, 0666)
	defer file.Close()
	if e != nil {
		fmt.Printf("创建[%s]异常:%v", PATH, e)
		return
	}
	writer := bufio.NewWriter(file)
	// 写入文件
	writer.WriteString(content)
	// 刷新内容
	writer.Flush()
}
