package main

import (
	"fmt"
	"strconv"
)

// select 用来同时等待多个 channel 的消息
// select 一次只会选择一个
func main() {
	c1 := make(chan string, 2)
	c2 := make(chan string, 2)

	go func() {
		for i := range 5 {
			c1 <- "c1 " + strconv.Itoa(i)
		}
	}()

	go func() {
		for i := range 5 {
			c2 <- "c2 " + strconv.Itoa(i)
		}
	}()

	for _ = range 10 {
		select {
		case m1 := <-c1:
			fmt.Println(m1)
		case m2 := <-c2:
			fmt.Println(m2)
		}
	}

}
