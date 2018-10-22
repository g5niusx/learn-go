package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
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
}
