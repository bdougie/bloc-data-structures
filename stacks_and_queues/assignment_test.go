package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStackIsInitializedEmpty(t *testing.T) {
	s := new(Stack)
	assert.Equal(t, true, s.IsEmpty())
}

func TestItemCanBePushedOntoStack(t *testing.T) {
	s := new(Stack)
	s.Push("12")
	assert.Equal(t, false, s.IsEmpty())
}

func TestItemCanBePopOffStack(t *testing.T) {
	s := new(Stack)
	s.Push("7")
	assert.Equal(t, false, s.IsEmpty())
	s.Pop()
	assert.Equal(t, true, s.IsEmpty())
}

func TestStackIsLIFO(t *testing.T) {
	s := new(Stack)
	s.Push("First")
	s.Push("Last")
	s.Pop()
	top := s.top.value
	assert.Equal(t, "First", top)
}

func TestQueueIsInitializedEmpty(t *testing.T) {
	q := new(Queue)
	assert.Equal(t, true, q.IsEmpty())
}

func TestItemCanBeQueued(t *testing.T) {
	q := new(Queue)
	q.Enqueue("12")
	assert.Equal(t, false, q.IsEmpty())
}

func TestItemCanBeDequeued(t *testing.T) {
	q := new(Queue)
	q.Enqueue("12")
	assert.Equal(t, false, q.IsEmpty())
	q.Dequeue()
	assert.Equal(t, true, q.IsEmpty())
}

func TestFIFO(t *testing.T) {
	q := new(Queue)
	q.Enqueue("First")
	q.Enqueue("Last")
	q.Dequeue()
	top := q.top.value
	assert.Equal(t, "Last", top)
}

func TestStackQueueCanHold2Stacks(t *testing.T) {
	one := new(Stack)
	two := new(Stack)
	q := StackQueue{one, two}
	assert.NotNil(t, q)
}

func TestStackQueueIsEmpty(t *testing.T) {
	one := new(Stack)
	two := new(Stack)
	q := StackQueue{one, two}
	assert.Equal(t, true, q.IsEmpty())
}

func TestStackQueueCanEnqueue(t *testing.T) {
	one := new(Stack)
	two := new(Stack)
	q := StackQueue{one, two}
	q.Enqueue("stuff")
	alphaValue := q.alpha.top.value
	assert.Equal(t, "stuff", alphaValue)
}

func TestStackQueueCanDequeueAndMaintainTheStackFIFO(t *testing.T) {
	one := new(Stack)
	two := new(Stack)
	q := StackQueue{one, two}
	q.Enqueue("stuff")
	q.Enqueue("more stuff")
	q.Dequeue()
	alphaValue := q.alpha.top.value
	assert.Equal(t, "more stuff", alphaValue)
}

func TestStackQueueCanPop(t *testing.T) {
	in := new(Stack)
	out := new(Stack)
	q := StackQueue{in, out}
	q.Enqueue("a thing")
	q.Enqueue("another thing")
	q.Dequeue()
	alphaSize := q.beta.size
	assert.Equal(t, 1, alphaSize)
}

func TestQueueStackIsEmpty(t *testing.T) {
	one := new(Queue)
	two := new(Queue)
	s := QueueStack{one, two}
	assert.Equal(t, true, s.IsEmpty())
}

func TestQueueStackCanPush(t *testing.T) {
	alpha := new(Queue)
	beta := new(Queue)
	s := QueueStack{alpha, beta}
	s.Push("YO")
	s.Push("YO")
	alphaSize := s.alpha.size
	assert.Equal(t, 2, alphaSize)
}

func TestQueueStackCanPop(t *testing.T) {
	alpha := new(Queue)
	beta := new(Queue)
	s := QueueStack{alpha, beta}
	s.Push("12")
	s.Push("3")
	s.Pop()
	alphaSize := s.alpha.size
	assert.Equal(t, 1, alphaSize)
}
