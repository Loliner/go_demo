package main

import (
	"fmt"
)

type Stack struct {
	items []int
}

func NewStack() *Stack {
	return &Stack{items: []int{}}
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	tmp := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return tmp, nil
}

func (s *Stack) Peek() (int, error) {
	if len(s.items) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack) Size() int {
	return len(s.items)
}

func exercise() {
	fmt.Println("=== Exercise: pointer ===")

	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	for stack.Size() > 0 {
		if value, err := stack.Pop(); err != nil {
			fmt.Printf("Pop error: %s\n", err)
		} else {
			fmt.Printf("Pop value %d\n", value)
		}
	}
}
