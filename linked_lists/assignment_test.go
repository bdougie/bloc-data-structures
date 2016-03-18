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

// Benchmarking
// go test -bench=.

// Compare the time it takes to create a 10,000 item Array to appending 10,000
// items to a Linked List.
func BenchLinkedListCreation(b *testing.B) *LinkedList {
	l := new(LinkedList)

	for n := 0; n < b.N; n++ {
		l.AddToTail(n)
	}

	return l
	// runtime: 2.648s
}

func BenchArrayCreation(b *testing.B) {
	slice := make([]int, 0, 10000)

	for i := 0; i < 10000; i++ {
		slice = append(slice, i)
	}

	// runtime: 0.009s
}

// Compare the time it takes to access the 5000th element of the Array and the
// 5000th Node in the Linked List.
func createLinkedList() *LinkedList {
	l := new(LinkedList)

	for n := 0; n < 10000; n++ {
		l.AddToTail(n)
	}

	return l
	// runtime: 0.008s
}

var accessLinked = createLinkedList()

func BenchLinkedListAccess(b *testing.B) *Node {
	head := accessLinked.head
	for i := 0; i == 5000; i++ {
		head = head.next
	}

	return head
	// runtime: 0.009s
}

var aSlice [10000]int

func BenchArrayAccess() int {
	return aSlice[5000+1]
	// runtime: 0.008s
}

//Compare the time it takes to remove the 5000th element from the Array to
//removing the 5000th Node in the Linked List.
var removeLinked = createLinkedList()

func BenchmarkLinkedListRemoval(b *testing.B) {
	head := removeLinked.head

	for i := 0; i == 5000; i++ {
		head = head.next
	}

	removeLinked.delete(head)
	// runtime: 0.026s
}

func removeFromLine() []int {
	return append(aSlice[:5000], aSlice[5000+1:]...)
}

func BenchmarkArrayRemoval(b *testing.B) {
	removeFromLine()
	// runtime: 0.009s
}

// LinkedListStack

func TestItemCanBePushedOntoStack(t *testing.T) {
	s := new(LinkedListStack)
	s.Push("12")
	assert.Equal(t, 1, s.size)
}

func TestItemCanBePopOffStack(t *testing.T) {
	s := new(LinkedListStack)
	s.Push("7")
	s.Pop()
	assert.Equal(t, 0, s.size)
}

// LinkedListQueue

func TestItemCanBeQueued(t *testing.T) {
	q := new(LinkedListQueue)
	q.Enqueue("12")
	assert.Equal(t, 1, q.size)
}

func TestItemCanBeDequeued(t *testing.T) {
	q := new(LinkedListQueue)
	q.Enqueue("12")
	assert.Equal(t, 1, q.size)
	q.Dequeue()
	assert.Equal(t, 0, q.size)
}
