package main

import "fmt"

func zeroByVal(val int) {
	val = 0
}

func zeroByPtr(iPtr *int) {
	*iPtr = 0
}

func main() {
	i := 1
	fmt.Println(i)
	zeroByVal(i)
	fmt.Println(i)

	j := 1
	fmt.Println(j)
	zeroByPtr(&j)
	fmt.Println(j)
}
