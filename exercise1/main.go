package main

import (
	"./stack"
	"fmt"
)

func main() {
	stack := stack.NewStack()

	stack.Push(1)
	stack.Push(2)
	stack.Pop()
	stack.Push(3)
	stack.Push(4)
	stack.Pop()
	stack.Push(5)

	for stack.Len() > 0 {
		fmt.Printf("%v\n", stack.Pop())
	}
}
