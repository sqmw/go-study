package main

import "fmt"

func MapKeys[k comparable, v interface{}](m map[k]v) []k {
	desKeys := make([]k, 0)
	for _k := range m {
		desKeys = append(desKeys, _k)
	}
	return desKeys
}

func main() {
	m := map[string]string{
		"1": "1",
		"2": "2",
	}
	fmt.Println(MapKeys[string, string](m))

	var list *List[int] = new(List[int])

	list.Push(1)
	list.Push(2)
	list.Push(3)
	list.Push(4)

	list.ShowAll()

}
