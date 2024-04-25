package main

import (
	"fmt"
	"maps"
)

func main() {
	// 直接申明赋值，此时不需要使用申请空间的关键字 make
	m := map[int]int{1: 100, 2: 2}
	val, in := m[3]
	fmt.Println(val, in)
	// 申明，然后赋值
	var m2 map[string]int
	fmt.Println(m2)
	m2 = make(map[string]int)
	m2["1"] = 1
	fmt.Println(m2)
	var m2C = map[string]int{"1": 1}
	fmt.Println(maps.Equal(m2, m2C))
	// clear 是专门针对 map 设计的函数
	clear(m2)
	fmt.Println(m2)
	m2["hhh"] = 123
	fmt.Println(m2)
	delete(m2, "hhh")
	fmt.Println(m2)
}
