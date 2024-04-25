package main

import (
	"fmt"
	"time"
)

func testTimeout1() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}
}

func testTimeout2() {
	c2 := make(chan int)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- 1
	}()
	select {
	case val := <-c2:
		fmt.Println("未超时", val)
	case <-time.After(time.Second * 3):
		fmt.Println("超时")
	}
}

func main() {
	testTimeout2()
}
