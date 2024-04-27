package main

import "fmt"

func main() {
	msgs := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-msgs:
		fmt.Println("收到msg: ", msg)
	default:
		fmt.Println("没有收到信息")
	}

	msg := "hi"
	select {
	case msgs <- msg:
		fmt.Println("发送消息")
	default:
		fmt.Println("没有发送消息")
	}

	select {
	case msg := <-msgs:
		fmt.Println("收到信息", msg)
	case sig := <-signals:
		fmt.Println("收到信号", sig)
	default:
		fmt.Println("no activity")
	}
}
