package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNodeIsInitialized(t *testing.T) {
	n := new(Node)
	assert.NotNil(t, n)
}

func TestLinkedListIsInitialized(t *testing.T) {
	l := new(LinkedList)
	assert.NotNil(t, l)
}

func TestLinkedListAddsToTailAndSetsHead(t *testing.T) {
	l := new(LinkedList)

	l.AddToTail("data things")
	assert.Equal(t, "data things", l.tail.value)
	assert.Equal(t, "data things", l.head.value)
}

func TestLinkedListAddsLatestTail(t *testing.T) {
	l := new(LinkedList)

	l.AddToTail("oranges")
	l.AddToTail("bananas")
	assert.Equal(t, "bananas", l.head.next.value)
}

func TestLinkedListRemovesLatestTail(t *testing.T) {
	l := new(LinkedList)

	l.AddToTail("oranges")
	l.AddToTail("bananas")
	l.RemoveTail()

	assert.Equal(t, "oranges", l.tail.value)
}

func TestRepresentationWasPrintedWithoutError(t *testing.T) {
	l := new(LinkedList)

	l.AddToTail("oranges")
	l.AddToTail("bananas")

	l.print()
}

func TestNodesDeleteFromLinkedList(t *testing.T) {
	l := new(LinkedList)

	l.AddToTail("oranges")
	l.AddToTail("apples")
	l.AddToTail("bananas")

	a := l.head.next
	b := l.tail
	l.delete(a)

	assert.Equal(t, b, l.head.next)
}

func TestNodesCanBeAddedToFront(t *testing.T) {
	l := new(LinkedList)
	n := Node{value: "oranges"}
	nr := Node{value: "bananas"}

	l.AddToTail(&n)
	l.AddToFront(&nr)

	assert.Equal(t, &nr, l.head)
}

func TestNodeCanBeRemovedFromHead(t *testing.T) {
	l := new(LinkedList)
	l.AddToTail("oranges")
	l.AddToTail("pizza")

	n := l.RemoveFront()

	assert.Equal(t, "oranges", n.value)
}
