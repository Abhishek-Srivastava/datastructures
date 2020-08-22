package linkedlist

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

func NewNode(val int) *Node {
	return &Node{Val: val}
}

func AddNode(head *Node, val int) *Node {
	if head == nil {
		head = NewNode(val)
		return head
	}
	temp := head
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = NewNode(val)
	return head
}

func Traverse(head *Node) {
	temp := head
	for temp != nil {
		fmt.Printf("%d ", temp.Val)
		temp = temp.Next
	}
	fmt.Println()
}

func MidPoint(head *Node) (*Node, int) {
	temp := head
	temp2 := head
	i := 1
	for temp.Next != nil && temp.Next.Next != nil {
		i++
		temp2 = temp2.Next
		temp = temp.Next.Next
	}
	return temp2, i
}

func Length(head *Node) int {
	temp := head
	i := 0
	for temp != nil {
		temp = temp.Next
		i++
	}
	return i
}

func reverse(head *Node) *Node {
	// 1 2 3 nil
	var prev *Node    // nil
	cur := head       // 1
	next := head.Next // 2
	for {
		cur.Next = prev // nil <1 <2 <3
		prev = cur      // 3
		if next == nil {
			return cur
		}
		cur = next
		next = next.Next
	}
}

/*
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	head := newNode(arr[0])
	for i := 1; i < len(arr); i++ {
		addNode(head, arr[i])
	}
	traverse(head)
	cur := reverse(head)
	traverse(cur)
}
*/
