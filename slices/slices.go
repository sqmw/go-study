package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println(s)
	s = append(s, "1")
	fmt.Println(s)

	s = make([]string, 3)
	fmt.Println(len(s), cap(s))

	s2 := s[1:]
	fmt.Println(len(s2), cap(s2))

	s2[0] = "100"

	fmt.Println(s, s2)

	var coppedDes []string
	copy(coppedDes, s)
	fmt.Println(coppedDes) // []

	// 下面的两个方法都是可以的
	//var coppedDes2 []string = []string{"", "", ""}
	coppedDes2 := make([]string, len(s))
	copy(coppedDes2, s)
	fmt.Println(coppedDes2)
	fmt.Println(slices.Equal(s, coppedDes))
}
