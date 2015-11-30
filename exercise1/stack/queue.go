package stack

import (
	"fmt"
)

// Queue implementation
type Queue struct {
	first *Node
	last  *Node
	size  int
}

// Base node object
type Node struct {
	value interface{}
	next  *Node
}

// Constructor
func NewQueue() *Queue {
	return &Queue{nil, nil, 0}
}

// Returns length of queue.
func (q *Queue) Len() int {
	return q.size
}

// Offer new Node.
func (q *Queue) Offer(value interface{}) {
	if q.size == 0 {
		q.first = &Node{value, nil}
		q.last = q.first
	} else {
		q.last.next = &Node{value, nil}
		q.last = q.last.next
	}
	q.size++
}

// Poll front Node.
// Return Node value.
func (q *Queue) Poll() (value interface{}) {
	if q.size > 0 {
		value = q.first.value
		q.first = q.first.next
		q.size--
		return
	} else {
		fmt.Println("Queue is empty!")
	}
	return
}

// Peek value of front Node in queue.
func (q *Queue) Peek() (value interface{}) {
	if q.size == 0 {
		fmt.Println("Queue is empty!")
		return
	}
	value = q.first.value
	return
}

// Check whether queue is empty.
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}
