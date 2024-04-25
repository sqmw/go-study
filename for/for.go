package main

import (
	"fmt"
)

// 因为Python解释的特性，因此在调用函数的时候需要已经扫描过这个函数，但是Go语言不一样，可以在后面
func main() {
	_for5()
}

func _for1() {
	i := 1
	for i <= 4 {
		fmt.Println(i)
		i += 1
	}
}

func _for2() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func _for3() {
	var i int = 0
	for i < 4 {
		fmt.Println(i)
		i++
	}
}

func _for4() {
	for i := range 4 {
		fmt.Println(i)
	}
	// bubble_sort
}

func _for5() {
	for i := range 100 {
		if i%2 == 0 {
			continue
		} else {
			fmt.Println(i)
		}
	}
}
