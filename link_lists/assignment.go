package main

import ()

type Node struct {
	value interface{}
	next  *Node
}

type LinkList struct {
	head, tail *Node
}

// http://l3x.github.io/golang-code-examples/2014/07/23/doubly-linked-list.html
// https://golang.org/pkg/container/list/
