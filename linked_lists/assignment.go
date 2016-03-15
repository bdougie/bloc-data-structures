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
