package main

import (
	"fmt"
	"time"
)

func main() {
	// 常规使用
	i := 4
	switch i {
	case 1:
		fmt.Println("i: ", 1)
	case 2:
		fmt.Println("i: ", 2)
	case 3, 4:
		fmt.Println(3, 4)
	default:
		fmt.Println("i: ", "default")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}

	_score := 59
	// 这里其实就是 if else
	switch {
	case _score >= 60:
		fmt.Println("pass")
	default:
		fmt.Println("try again")
	}

}
