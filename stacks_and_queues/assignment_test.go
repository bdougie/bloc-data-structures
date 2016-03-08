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
	inboxValue := q.inbox.top.value
	assert.Equal(t, "stuff", inboxValue)
}

func TestStackQueueCanDequeueAndMaintainTheStackLIFO(t *testing.T) {
	one := new(Stack)
	two := new(Stack)
	q := StackQueue{one, two}
	q.Enqueue("stuff")
	q.Enqueue("more stuff")
	q.Dequeue()
	inboxValue := q.inbox.top.value
	assert.Equal(t, "stuff", inboxValue)
}

func TestStackQueuePopulatesOutboxWhenOutBoxIsEmpty(t *testing.T) {
	in := new(Stack)
	out := new(Stack)
	q := StackQueue{in, out}
	q.Enqueue("a thing")
	q.Enqueue("another thing")
	q.Dequeue()
	outboxValue := q.outbox.top.value
	assert.Equal(t, "another thing", outboxValue)
}

func TestQueueStackIsEmpty(t *testing.T) {
	one := new(Queue)
	two := new(Queue)
	s := QueueStack{one, two}
	assert.Equal(t, true, s.IsEmpty())
}

func TestQueueStackCanPush(t *testing.T) {
	inbox := new(Queue)
	outbox := new(Queue)
	s := QueueStack{inbox, outbox}
	s.Push("YO")
	assert.Equal(t, false, s.inbox.IsEmpty())
}

func TestQueueStackCanPop(t *testing.T) {
	inbox := new(Queue)
	outbox := new(Queue)
	s := QueueStack{inbox, outbox}
	s.Push("YO")
	s.Push("YO")
	s.Pop()
	outboxValue := s.outbox.top.value
	assert.Equal(t, "YO", outboxValue)
}
