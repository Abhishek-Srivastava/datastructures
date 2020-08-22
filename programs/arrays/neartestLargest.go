package arrays

import (
	"datastructures/programs/stack"
	"fmt"
	"math"
)

var maxStack *stack.Stack
var largerIndex []int

func nearestLargerRight(a []int) {
	max := math.MinInt32
	largerIndex = append(largerIndex, max)
	maxStack.Push(a[len(a)-1])
	for i := len(a) - 2; i >= 0; i-- {
		if a[i] <= maxStack.Top() {
			largerIndex = append(largerIndex, maxStack.Top())
			maxStack.Push(a[i])
		} else {
			for {
				maxStack.Pop()
				if maxStack.Top() == -1 {
					largerIndex = append(largerIndex, max)
					break
				}
				if a[i] <= maxStack.Top() {
					largerIndex = append(largerIndex, maxStack.Top())
					maxStack.Push(a[i])
					break
				}
			}
		}
	}
}

func Run() {
	maxStack = stack.New()
	a := []int{20, 15, 10, 3, 5, 17, 18}
	nearestLargerRight(a)

	for i := 0; i < len(largerIndex)/2; i++ {
		largerIndex[i], largerIndex[len(largerIndex)-i-1] = largerIndex[len(largerIndex)-i-1], largerIndex[i]
	}
	fmt.Println(a)
	fmt.Println(largerIndex)
}
