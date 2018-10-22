package main

import (
	"fmt"
	"sort"
)

func main() {
	pecker := Pecker{}
	fmt.Printf("啄木鸟的飞行速度是:%vkm/h\n", pecker.fly())

	var slice HeroSlice = make([]Hero, 0)
	slice = append(slice, Hero{name: "德玛西亚", hurt: 20, speed: 40}, Hero{name: "刀锋意志", hurt: 40, speed: 50}, Hero{name: "无极剑圣", hurt: 100, speed: 100})
	sort.Sort(slice)
	fmt.Printf("根据速度排序以后:%v\n", slice)
}

type Fly interface {
	fly() int
}

type Pecker struct {
}

func (pecker Pecker) fly() int {
	return 50
}

type Hero struct {
	name  string
	hurt  int
	speed int
}

type HeroSlice []Hero

func (slice HeroSlice) Len() int {
	return len(slice)
}

// 该方法决定使用升序还是降序
func (slice HeroSlice) Less(i, j int) bool {
	return slice[i].speed > slice[j].speed
}

func (slice HeroSlice) Swap(i, j int) {
	temp := slice[i]
	slice[i] = slice[j]
	slice[j] = temp
}
