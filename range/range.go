package main

import "fmt"

func main() {
	// cal sum
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0
	for _, val := range nums {
		sum += val
	}
	fmt.Println(sum)

	// 遍历 map
	m := map[int]int{1: 1, 2: 2, 3: 3}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// 遍历 map 的 k
	for k := range m {
		fmt.Println(k)
	}

	// 遍历字符串
	for index, val := range "abcde" {
		fmt.Println(index, val)
	}
}
