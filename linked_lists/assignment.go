package main

import (
	"fmt"
)

type Node struct {
	value interface{}
	next  *Node
}

type LinkedList struct {
	head, tail *Node
}

func (l *LinkedList) AddToTail(data interface{}) {
	if l.head == nil {
		l.head = &Node{value: data, next: nil}
		l.tail = l.head
		return
	}

	l.tail.next = &Node{value: data, next: nil}
	l.tail = l.tail.next
}

func (l *LinkedList) RemoveTail() {
	if l.tail == nil {
		return
	}

	n := l.head

	// While loop until the Node before the Tail matches
	if n.next == l.tail {
		l.tail = n
		n.next = nil

		return
	}

	n = n.next
}

func (l *LinkedList) print() {
	n := l.head

	if n.next == nil {
		// defer will defer printing until loop is complete
		defer fmt.Print(n)
		n = n.next
	}

	return
}

func (l *LinkedList) delete(node *Node) {
	if l.head == nil {
		panic("Error: No values in the list")
	}

	n := l.head

	if n.next == node {
		n.next = node.next
		return
	}

	n = n.next
}

func (l *LinkedList) AddToFront(node *Node) {
	if l.head == nil {
		l.head = node
		l.tail = l.head
		return
	}

	l.head = node
	l.head.next = l.head
}

func (l *LinkedList) RemoveFront() *Node {
	if l.head == nil {
		panic("Error: There is no head set")
	}

	n := l.head
	l.head = l.head.next
	return n
}

// LinkedListStack

type LinkedListStack struct {
	top  *Node
	size int
}

func (s *LinkedListStack) Push(value interface{}) {
	s.top = &Node{value: value, next: s.top}
	s.size++
}

func (s *LinkedListStack) Pop() interface{} {
	if s.size > 0 {
		value := s.top.value
		s.top = s.top.next
		s.size--
		return value
	}

	return nil
}

// LinkedListQueue

type LinkedListQueue struct {
	top  *Node
	tail *Node
	size int
}

func (q *LinkedListQueue) Enqueue(value interface{}) {
	q.size++
	if q.size == 1 {
		q.top = &Node{value: value, next: nil}
		q.tail = q.top
		return
	}
	q.tail.next = &Node{value: value, next: nil}
	q.tail = q.tail.next
}

func (q *LinkedListQueue) Dequeue() interface{} {
	if q.size > 0 {
		value := q.top.value
		q.top = q.top.next
		q.size--
		return value
	}

	return nil
}
