package queues

import (
	"datastructures/programs/stack"
	"fmt"
)

func queue(val int) {
	s1.Push(val)
	printQueue()
}

func printQueue() {
	if len(s1.Val) > 0 {
		fmt.Println(s1.Val)
	}
	for i := len(s2.Val) - 1; i >= 0; i-- {
		fmt.Printf("%d ", s2.Val[i])
	}
	fmt.Println()
}

func dequeue() int {

	for {
		if len(s1.Val) > 0 {
			topItem, _ := s1.Pop()
			s2.Push(topItem)
		} else {
			if len(s2.Val) > 0 {
				poppedEl, _ := s2.Pop()
				printQueue()
				return poppedEl
			}
			panic(fmt.Errorf("Queue empty"))
		}
	}
}

var s1 *stack.Stack
var s2 *stack.Stack

func Run() {
	s1 = &stack.Stack{
		Val: []int{},
	}
	s2 = &stack.Stack{
		Val: []int{},
	}
	// 1 2 3 4
	queue(1)
	queue(2)
	queue(3)
	queue(4)
	fmt.Printf("Dequeued: %d\n", dequeue())
	fmt.Printf("Dequeued: %d\n", dequeue())
	fmt.Printf("Dequeued: %d\n", dequeue())
	fmt.Printf("Dequeued: %d\n", dequeue())
	queue(5)
	queue(6)
	queue(7)
	fmt.Printf("Dequeued: %d\n", dequeue())
	printQueue()
}
