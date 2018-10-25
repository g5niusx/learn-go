package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "", "用户名称")
	flag.Parse()
	fmt.Printf("输入名称为:%s", name)
}
