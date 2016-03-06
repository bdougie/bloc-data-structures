package main

import ()

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value interface{}
	stack *Stack
	next  *Element
}

func (s *Stack) isEmpty() bool {
	if s.size > 0 {
		return false
	}

	return true
}

func (s *Stack) Push(value interface{}) {
	s.top = &Element{value: value, stack: s, next: s.top}
	s.size++
}

func (s *Stack) Pop() interface{} {
	if s.size > 0 {
		value := s.top.value
		s.top = s.top.next
		s.size--
		return value
	}

	return nil
}
