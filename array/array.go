package main

import "fmt"

func main() {
	// 数字的类型和Java的还是一样的
	var arr1 [5]int
	fmt.Println(arr1)
	fmt.Println("len: ", len(arr1))

	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl", arr2)

	// 这个是切片, 不是数组，Python里面的切片，其实就是将数组的某一部分进行浅拷贝到另外一个数组
	arr3 := []int{1, 2, 3, 4, 5}
	fmt.Println(arr3)

	arr4 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr4)

	// 这个表示 index >= 6 的设置为 4,5
	// 其中 4: index == 6; 5: index == 7
	arr5 := [...]int{1, 6: 4, 5}
	fmt.Println(arr5)
}
