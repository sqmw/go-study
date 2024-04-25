package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r *rect) area() float64 {
	return r.height * r.width
}

func (r *rect) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.perimeter())
	fmt.Println(g.area())
}

func main() {
	c := circle{1}
	r := rect{1, 2}

	measure(&c)
	measure(&r)
}
