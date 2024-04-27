package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)
	done := make(chan bool)
	go func() {
		for {
			select {
			// 表示需要结束了
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("<-ticker.C", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
}
