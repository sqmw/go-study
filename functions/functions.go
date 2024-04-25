package main

import (
	"fmt"
	"reflect"
)

// 单个返回值的
func flus(a, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

// 多个返回值的
func multiReturn() (int, int) {
	return 1, 2
}

// 接收的参数是 slices
func plusSlices(nums []int) int {
	_sum := 0
	for _, num := range nums {
		_sum += num
	}
	return _sum
}

// 参数可变
func plusVariadicFunction(nums ...int) int {
	// 本质上这个就是一个 slices
	fmt.Println(reflect.TypeOf(nums))
	_sum := 0
	for _, num := range nums {
		_sum += num
	}
	return _sum
}

// closures
// 闭包其实就是让一个变量来接受一个函数作为参数，然后调用这个参数，这个参数具有外层包围的内层的所有的数据
func nextInt() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// 通过闭包实现 慢斐波拉契数列
func fibClosure(n int) int {
	var slowFib func(n int) int
	slowFib = func(n int) int {
		// 这里使用的 n 符合最短匹配原则
		if n <= 1 {
			return n
		} else {

			// 必须申明，不申明这里直接报错
			return slowFib(n-1) + slowFib(n-2)
		}
	}
	return slowFib(n)
}
func main() {
	// 单个返回值
	//fmt.Println(flus(1, 100))
	//fmt.Println(plusPlus(1, 2, 3))
	// 多个返回值
	//fmt.Println(multiReturn())
	//r1, r2 := multiReturn()
	//fmt.Println(r1, r2)
	//fmt.Println(plusSlices([]int{1, 2, 3, 4}))
	//fmt.Println(plusVariadicFunction([]int{1, 2, 3}...))
	//nextInt := nextInt()
	//fmt.Println(nextInt())
	fmt.Println(fibClosure(45))
}
