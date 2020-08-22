package stack

import (
	"fmt"
)

type Stack struct {
	Val    []int
	minVal int
}

func New() *Stack {
	return &Stack{
		Val: []int{},
	}
}

func (s *Stack) Top() int {
	if len(s.Val) > 0 {
		return s.Val[len(s.Val)-1]
	}
	return -1
}

func (s *Stack) Push(val int) {
	s.Val = append(s.Val, val)
}

func (s *Stack) Pop() int {
	if len(s.Val) == 0 {
		return -1
	}
	poppedEle := s.Val[len(s.Val)-1]
	s.Val = s.Val[:len(s.Val)-1]
	return poppedEle
}

/*
func (s *Stack) Push(val int) {
	if len(s.Val) == 0 {
		s.Val = append(s.Val, val)
		s.minVal = val
	} else {
		if val >= s.minVal {
			s.Val = append(s.Val, val)
		} else {
			funcVal := 2*val - s.minVal
			// fmt.Println("Appending : ", funcVal)
			s.Val = append(s.Val, funcVal)
			s.minVal = val
		}
	}
	// fmt.Println("minVal, value being pushed", s.minVal, val)
}

func (s *Stack) Pop() (int, error) {
	if len(s.Val) == 0 {
		return -1, (fmt.Errorf("Stack empty"))
	}
	mostRecent := s.Val[len(s.Val)-1]
	if mostRecent > s.minVal {
		s.Val = s.Val[:len(s.Val)-1]
		return mostRecent, nil
	} else {
		k := s.minVal
		s.minVal = 2*s.minVal - mostRecent
		s.Val = s.Val[:len(s.Val)-1]
		return k, nil
	}
}

func (s *Stack) reverse() {
	s.solve(s.Val)
}

func (s *Stack) solve(a []int) {
	if len(s.Val) == 0 {
		return
	}
	temp, _ := s.Pop()
	fmt.Println("popping", temp)
	s.solve(a)
	fmt.Println("inserting", temp)
	s.insert(temp)
}

func (s *Stack) insert(item int) {
	if len(s.Val) == 0 {
		fmt.Println("inserting finally", item)
		s.Push(item)
		fmt.Println("After inserting the stack is", s.Val)
		return
	}
	temp, _ := s.Pop()
	fmt.Println("popping during insertion", temp)
	s.insert(item)
	s.Push(temp)
	fmt.Println("Now pushing post insertion", temp, s.Val)
}
*/

func Run() {
	arr := []int{1, 2, 3, 4, 5, 6}
	s := &Stack{
		Val: []int{},
	}
	for _, n := range arr {
		s.Push(n)
	}
	fmt.Println(s.Val)
	//s.reverse()
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Val)
	// fmt.Println(s.Pop())
	// fmt.Println(s.Pop())
	// fmt.Println(s.Pop())
}
