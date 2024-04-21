package main

import "fmt"

// 变量的命名遵从 小驼峰，和Java一样的

// 使用 var 的时候，一个 var 申明多个 变量的时候，变量必须是同一种类型的，
// 这样也符合自然语言的阅读、逻辑(因为我使用Python来引入nonlocal的时候就是这么使用的)
var globalNum int = 1

func main() {
	s := "_string"
	fmt.Println(globalNum)
	fmt.Println(s)
	var i, j int = 1, 2
	fmt.Println(i, j)
	t := "temp"
	fmt.Println(t)
}
