package main

import ()

// Stack
//

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

// Queue
//

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

// StackQueue
//

type StackQueue struct {
	alpha *Stack
	beta  *Stack
}

func (q *StackQueue) IsEmpty() bool {
	if q.alpha.IsEmpty() && q.beta.IsEmpty() {
		return true
	}

	return false
}

func (q *StackQueue) Enqueue(value interface{}) {
	// move all items to beta
	if !q.alpha.IsEmpty() {
		q.beta.Push(q.alpha.Pop())
	}

	// push new value to aplha
	q.alpha.Push(value)

	// move all items back from beta to alpha
	if !q.beta.IsEmpty() {
		q.alpha.Push(q.beta.Pop())
	}
}

func (q *StackQueue) Dequeue() interface{} {
	if !q.alpha.IsEmpty() {
		q.beta.Push(q.alpha.Pop())
	}

	return nil
}

// QueueStack
//

type QueueStack struct {
	alpha *Queue
	beta  *Queue
}

func (s *QueueStack) IsEmpty() bool {
	if s.alpha.IsEmpty() && s.beta.IsEmpty() {
		return true
	}

	return false
}

func (s *QueueStack) Push(value interface{}) {
	s.alpha.Enqueue(value)
}

func (s *QueueStack) Pop() interface{} {
	// dequeue all elements to the beta queue
	if s.alpha.size > 1 {
		s.alpha.Enqueue(s.alpha.Dequeue())
	}

	// remove the remaining element
	s.alpha.Dequeue()

	// dequeue beta back to alpha
	s.alpha.Enqueue(s.beta.Dequeue())

	if s.IsEmpty() {
		return nil
	}

	return s.alpha.Dequeue()
}
