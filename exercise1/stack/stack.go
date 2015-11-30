package stack

import (
	"fmt"
)

// Basic concept:
// queue1 will handle all pushes.
// If there is a pop, poll all nodes from queue1 and offer them into queue2 except last node in queue1.
// Return the last node value as the pop value.
// Switch names of queue1 and queue2 after each pop.
type Stack struct {
	queue1 *Queue
	queue2 *Queue
}

// Constructor
func NewStack() *Stack {
	return &Stack{NewQueue(), NewQueue()}
}

// Return length of stack
func (s *Stack) Len() int {
	return s.queue1.Len() + s.queue2.Len()
}

// Push node into stack
func (s *Stack) Push(value interface{}) {
	s.queue1.Offer(value)
}

// Pop top node
func (s *Stack) Pop() (value interface{}) {
	if s.Len() == 0 {
		fmt.Println("Stack is Empty!")
		return
	}
	for s.queue1.Len() > 1 {
		s.queue2.Offer(s.queue1.Poll())
	}
	value = s.queue1.Poll()
	temp := s.queue1
	s.queue1 = s.queue2
	s.queue2 = temp
	return
}

// Peek the top node value
func (s *Stack) Peek() (value interface{}) {
	if s.Len() == 0 {
		fmt.Println("Stack is Empty!")
		return nil
	}
	for s.queue1.Len() > 1 {
		s.queue2.Offer(s.queue1.Poll())
	}
	value = s.queue1.Peek()
	return
}

// Check whether stack is empty
func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}
