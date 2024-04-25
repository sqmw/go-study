package main

import "fmt"

// 其实就是 in, in 这个参数可以是没有限制的那种
func sendOnly(in chan<- int, val int) {
	in <- val
}

// 将 out 的数据移动一个到 in 里面
func sendEitherReceive(in chan<- int, out <-chan int) {
	val := <-out
	in <- val
}

func main() {
	send := make(chan int, 1)
	receive := make(chan int, 1)

	sendOnly(send, 99)
	sendEitherReceive(receive, send)

	fmt.Println(<-receive)
}
