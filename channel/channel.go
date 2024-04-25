package main

// channel 可以申明只能 send(入) 或者 receive(出)，可以在申明的时候就限制，也可以在封装在函数申明参数的时候限制

import (
	"fmt"
	"time"
)

func ChannelWithoutBuff() {
	var chanel = make(chan any, 10)
	go func() {
		// 如果我们没有一个给chanel设置 buffer，那么就只能在 receiver 的时候同时进行 send
		chanel <- 1
	}()
	msg := <-chanel
	fmt.Println(msg)
}

func ChannelWithBuff() {
	// 队列的数学逻辑模型
	channel := make(chan any, 10)
	channel <- 1
	channel <- 2
	channel <- 3
	channel <- 4
	for val := range channel {
		fmt.Println(val)
	}
}

func ChannelSync() {
	// 嵌套定义方法
	var work func(done chan bool)
	work = func(done chan bool) {
		time.Sleep(time.Second * 2)
		done <- true
	}
	done := make(chan bool)
	// 没有 go 就会死锁，一个线程也会死锁
	go work(done)
	<-done
	fmt.Println("sync finish")
}

func main() {
	channel := make(chan int, 3)
	channel <- 1
	channel <- 2
	channel <- 3
	for _ = range 3 {
		fmt.Println(<-channel)
	}
}
