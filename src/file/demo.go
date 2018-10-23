package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"unicode"
)

// 统计文件里面的字数z
func main() {

	file, e := os.Open("/Users/g5niusx/Workspace/learn-go/src/file/main.go")
	if e != nil {
		fmt.Printf("打开文件异常:%v", e)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	charCount := CharCount{}
	for {
		s, i := reader.ReadString('\n')
		if i == io.EOF {
			break
		}
		if i != nil {
			fmt.Printf("读取文件异常:%v", e)
			break
		}
		// 使用rune切片。来支持中文
		for _, v := range []rune(s) {
			if v >= 'a' && v <= 'c' || v >= 'A' && v <= 'Z' {
				charCount.EnCount++
			} else if v == ' ' || v == '\t' {
				charCount.SpaceCount++
			} else if IsChineseChar(v) {
				charCount.CnCount++
			} else {
				charCount.OtherCount++
			}
		}
	}
	bytes, _ := json.Marshal(charCount)
	fmt.Printf("统计结果为:%v\n", string(bytes))
}

// 判断字符串是否是汉字
func IsChineseChar(c rune) bool {
	if unicode.Is(unicode.Scripts["Han"], c) {
		return true
	}
	return false
}

type CharCount struct {
	// 中文字数
	CnCount int64 `json:"cnCount"`
	// 英文字数
	EnCount int64 `json:"enCount"`
	// 空格字数
	SpaceCount int64 `json:"spaceCount"`
	// 其他字数
	OtherCount int64 `json:"otherCount"`
}
