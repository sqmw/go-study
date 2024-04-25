package main

import "fmt"

func main() {
	var chanel = make(chan any)
	go func() {
		chanel <- 1
	}()
	msg := <-chanel
	fmt.Println(msg)
}
