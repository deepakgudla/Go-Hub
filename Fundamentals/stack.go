package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	items []interface{}
}

type Stack_ds struct{}

func (s Stack_ds) Name() string {
	return "stack"
}

// Push
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Stack is empty")
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, nil
}

func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Stack is empty")
	}

	index := len(s.items) - 1
	item := s.items[index]
	return item, nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s Stack_ds) Run() {
	stack := &Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	top, err := stack.Peek()
	if err == nil {
		fmt.Println("Top element:", top)
	}

	for !stack.IsEmpty() {
		item, _ := stack.Pop()
		fmt.Println("Popped:", item)
	}
}

func init() {
	Register(Stack_ds{})
}
