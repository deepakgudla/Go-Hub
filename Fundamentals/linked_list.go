package main

import "fmt"

type node struct {
	data int
	next *node //pointing to the next element
}

type linkedlist struct {
	head   *node
	length int
}

type LinkedList_ds struct{}

func (l *linkedlist) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l LinkedList_ds) Name() string {
	return "Linked_list"
}

func (l LinkedList_ds) Run() {

	list := linkedlist{}

	node1 := &node{data: 13}
	node2 := &node{data: 77}
	node3 := &node{data: 3}
	node4 := &node{data: 5}
	node5 := &node{data: 35}

	list.prepend(node1)
	list.prepend(node2)
	list.prepend(node3)
	list.prepend(node4)
	list.prepend((node5))

	fmt.Println(list)
}

func init() {
	Register(LogicalOp{})
}
