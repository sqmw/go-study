package main

import (
	"fmt"
	"math"
)

const s string = "constant"

// 实际上就是 int32
const c rune = '爱'
const n = 50000

// 对于数字类型的 const, 只要没有显示申明的时候，是没有确定的类型的
// 1. 我们可以对其进行类型强转
// 2. 当我们需要调用函数使用这个 const 的时候，会自己适配对应的数据类型，前提是本来就是

func main() {
	fmt.Println(s)
	fmt.Printf("c: %d \n", c)
	const d = 3e20
	const f = 1.2
	fmt.Println(d / n)
	fmt.Println(math.Sin(n))
	// 直接报错
	fmt.Println(add(f, d))
}

func add(a int, b int) int {
	return a + b
}
