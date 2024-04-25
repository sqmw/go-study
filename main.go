package main

import "fmt"

type a struct {
	v int
}

func (_a *a) A() string {
	return fmt.Sprintf("$%d$", _a.v)
}

func main() {
	a := a{100}
	fmt.Println(a)
}
