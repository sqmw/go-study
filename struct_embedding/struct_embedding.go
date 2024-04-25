package main

import (
	"fmt"
	"reflect"
)

type base struct {
	num int
}

func (b *base) describe() string {
	return fmt.Sprint("base.num", b.num)
}

func (b base) describe2() string {
	return fmt.Sprint("2 base.num", b.num)
}

// container 其实就是 base 的孩子，继承的时候把 base 的接口也继承了
type container struct {
	base
	info string
}

// 接口里面定义的都是方法
type describer interface {
	describe() string
	describe2() string
}

func main() {
	b := base{
		num: 100,
	}

	fmt.Println(b)

	c := container{
		base: b,
		info: "child",
	}

	fmt.Println(reflect.TypeOf(c))

	var c2 describer = &c

	fmt.Println(c2.describe())
	fmt.Println(reflect.TypeOf(c2))
}
