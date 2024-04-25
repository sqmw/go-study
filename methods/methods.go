package main

import "fmt"

type rect struct {
	w, h int
}

// 意思是调用这个方法的时候需要使用指针
func (r *rect) perimeter() int {
	return 2 * (r.h + r.w)
}

// 意思是调用这个方法的时候需要使用 指针指向的对象
func (r rect) area() int {
	return r.h * r.w
}

// 函数是和 class 或者是 struct 没有关系的，但是 method 则是属于 class 或者是 struct
func main() {
	r := rect{w: 1, h: 2}
	fmt.Println(r.area())
	fmt.Println(r.perimeter())
}
