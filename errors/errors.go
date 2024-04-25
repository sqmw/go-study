package main

import (
	"errors"
	"fmt"
)

func f(v int) (int, error) {
	if v == 0 {
		return -1, errors.New("can't handle val 0")
	} else {
		return v + 3, nil
	}
}

var noMaterialError = fmt.Errorf("no material error")

func makeTea(material int) error {
	if material < 1 {
		return noMaterialError
	} else {
		fmt.Println("tea made")
		return nil
	}
}

func main() {
	for i := range 2 {
		if val, err := f(i); err == nil {
			fmt.Println(val)
		} else {
			fmt.Println(err)
		}
	}

	for i := range 2 {
		if err := makeTea(i); err != nil {
			fmt.Println(err)
		}
	}
}
