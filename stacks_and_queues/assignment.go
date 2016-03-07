package main

import ()

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value    interface{}
	stack    *Stack
	next     *Element
	previous *Element
}

func (s *Stack) IsEmpty() bool {
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

type Queue struct {
	top  *Element
	tail *Element
	size int
}

func (q *Queue) IsEmpty() bool {
	if q.size > 0 {
		return false
	}

	return true
}

func (q *Queue) Enqueue(value interface{}) {
	q.size++
	if q.size == 1 {
		q.top = &Element{value: value, previous: nil}
		q.tail = q.top
		return
	}
	q.tail.previous = &Element{value: value, previous: nil}
	q.tail = q.tail.previous
}

func (q *Queue) Dequeue() interface{} {
	if !q.IsEmpty() {
		value := q.top.value
		q.top = q.top.previous
		q.size--
		return value
	}

	return nil
}

type StackQueue struct {
	inbox  *Stack
	outbox *Stack
}

func (q *StackQueue) IsEmpty() bool {
	if q.inbox.IsEmpty() && q.outbox.IsEmpty() {
		return true
	}

	return false
}

func (q *StackQueue) Enqueue(value interface{}) {
	q.inbox.Push(value)
}

func (q *StackQueue) resetInbox() {
	for q.inbox.IsEmpty() {
		q.outbox.Push(q.inbox.Pop())
	}
}

func (q *StackQueue) Dequeue() interface{} {
	if !q.IsEmpty() {
		q.outbox.Push(q.inbox.Pop())
	}

	return nil
}

// type QueueStack struct {
// 	inbox  *Queue
// 	outbox *Queue
// }

// func (s *StackQueue) IsEmpty() bool {
// 	if s.inbox.IsEmpty() && s.outbox.IsEmpty() {
// 		return true
// 	}

// 	return false
// }

// func (s *Queue) Push(value interface{}) {
// 	s.top = &Element{value: value, stack: s, next: s.top}
// 	s.size++
// }

// func (s *StackQueue Pop() interface{} {
// 	if s.size > 0 {
// 		value := s.top.value
// 		s.top = s.top.next
// 		s.size--
// 		return value
// 	}

// 	return nil
// }
