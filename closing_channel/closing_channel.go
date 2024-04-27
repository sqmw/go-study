package main

import "fmt"

// channel 的关闭仅仅是 send 而已，对于已经在 channel 里面的数据，是可以进行接收的
// close(chanA) // 表示关闭 channel
// val, more := <- chanA
// more == False <==> channel 关闭&&为空，意思是表示的是可不可能有更多的数据

func channelClose() {
	c := make(chan int, 1)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		// close 之后使用 range 遍历的时候就不会报错了
		close(c)
	}()

	for val := range c {
		fmt.Println(val)
	}
}

func workHandle() {
	jobs := make(chan int, 5)
	// 没有申请buff的chan可以用来实现同步
	done := make(chan bool)

	go func() {
		for {
			if job, more := <-jobs; more {
				fmt.Println("收到了 job", job)
			} else {
				fmt.Println("已完成所有任务")
				done <- true
				break
			}
		}
	}()

	for i := range 4 {
		jobs <- i
	}
	close(jobs)

	<-done

	// 测试 more 为 False 的条件
	_, more := <-jobs
	fmt.Println("received more jobs:", more)
}
func main() {
	workHandle()
}
