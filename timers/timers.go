package main

import (
	"fmt"
	"time"
)

func main() {
	// 演示 time.NewTimer 类似 time.Sleep() 的使用
	t1 := time.NewTimer(time.Second)

	<-t1.C
	fmt.Println("t1 fired")
	// 终止 timer.NewTimer
	t2 := time.NewTimer(time.Second)
	go func() {
		<-t2.C
		fmt.Println("t2 fired")
	}()
	if stopped := t2.Stop(); stopped {
		fmt.Println("t2 stopped")
	}
}
