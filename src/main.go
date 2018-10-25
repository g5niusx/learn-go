package main

import "fmt"

func main() {
	str := "去他娘的1024程序员节"
	for _, c := range []rune(str) {
		fmt.Print(c, " ")
	}
}
