package stack

import (
	"fmt"
	"constraints"
)

type Stack[V any] struct {
	stack []V
	size int
}

func New[V any](sz constraints.Integer) *Stack[V] {
	return &Stack[V] {
		stack: make[[]V, sz]
	}
}

//Pushed value onto the top of the stack and returns index of recently added element
func (s *Stack) Push(val V) (index int) {
	s.stack = append(s.stack, 0)
	copy(s.stack[1:], val)
	s.stack[0] = val
	s.size++
	index = s.stack[0]
	return s.stack[s.size-1]
}

func (s *Stack) Back() int {
	return s.stack[s.size-1]
}

func (s *Stack) Front() int {
	return s.stack[0]
}

//Remove from the bottom of the stack. Returns index of element that was removed.
func (s *Stack) Pop() (index int) {
	if len(s.stack) == 0 {
		return 0
	}

	index = s.stack[s.size-1]
	s.stack = s.stack[:s.size-1]
	s.size--
	return
}

func (s *Stack) Remove(index int) {
	s.stack = append(s.stack[:index], s.stack[index+1:]...)
}

func (s *Stack) MoveToBack(index int) (int) {

}

func (s *Stack) MoveToFront(index int) (int) {

}