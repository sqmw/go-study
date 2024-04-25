package main

import (
	"errors"
	"fmt"
	"strconv"
)

type argError struct {
	arg int
	msg string
}

// 这里使用指针防止浪费内存
func (_argError *argError) Error() string {
	return fmt.Sprintf("%d-%s", _argError.arg, _argError.msg)
}

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with " + strconv.Itoa(arg)}
	}
	return arg + 3, nil
}

func main() {
	_, err := f(42)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae.msg)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
