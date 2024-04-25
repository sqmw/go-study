package main

import (
	"fmt"
	"time"
)

func printFiveTimes(val string) {
	for _, i := range [5]int{0, 1, 2, 3, 4} {
		fmt.Println(val, i)
	}
}

// goroutine 是一个轻量级的线程
func main() {
	// 自执行函数
	func() {
		printFiveTimes("Jack")
	}()

	go printFiveTimes("Tom")

	go printFiveTimes("Tim")

	time.Sleep(time.Second)
}
