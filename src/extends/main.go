package main

import "fmt"

func main() {
	undergraduate := Undergraduate{}
	undergraduate.Name = "tomcat"
	undergraduate.Student.Name = "测试"
	undergraduate.Age = 23
	undergraduate.Score = 90.0
	fmt.Printf("大学生的信息:%v\n", undergraduate)
	undergraduate.pass()
}

type Student struct {
	Name  string
	Age   int
	Score float64
}

func (student Student) pass() {
	fmt.Println("学生的考试通过了...😄")
}

func (undergraduate Undergraduate) pass() {
	fmt.Println("大学生的考试通过了...🐶")
}

type Undergraduate struct {
	Student
}
