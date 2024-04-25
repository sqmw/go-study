package main

import "fmt"

// ListNode 使用 go 实现一个 List
type ListNode[T any] struct {
	val  T
	next *ListNode[T]
}

type List[T any] struct {
	// ListNode 的类型是 ListNode[T]
	head, tail *ListNode[T]
}

func (list *List[T]) Push(v T) {
	// 表示是第一次插入
	if list.head == nil {
		list.head = &ListNode[T]{
			val: v,
		}
		list.tail = list.head
	} else {
		// 不是第一次插入
		list.tail.next = &ListNode[T]{val: v}
		list.tail = list.tail.next
	}
}

func (list *List[T]) ShowAll() {
	for p := list.head; p != nil; p = p.next {
		fmt.Println(p.val)
	}
}
