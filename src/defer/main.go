package main

import "fmt"

func main() {
	defer release()
	fmt.Println("做一些数据库操作..")
	a := 1
	b := 2
	test(&a, &b)
}

func release() {
	fmt.Println("释放资源....")
}
func test(a *int, b *int) int {
	defer fmt.Printf("释放变量a %v \n", *a)
	defer fmt.Printf("释放变量b %v \n", *b)
	// defer会对变量进行值拷贝,所以a++和b++并不会影响a和b的输出,在这儿使用了指针来测试。可以发现指针也被拷贝了，不会影响原来的值
	*a++
	*b++
	return *a * *b
}
