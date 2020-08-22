package sorting

import (
	ll "datastructures/programs/linkedlist"
)

// MergeSort se
func MergeSort(head *ll.Node) *ll.Node {
	if head == nil || head.Next == nil {
		return head
	}
	mp, _ := ll.MidPoint(head)
	nextToMid := mp.Next
	mp.Next = nil
	lh := MergeSort(head)
	rh := MergeSort(nextToMid)
	return mergeSortedList(lh, rh)
}

func mergeSortedList(lh, rh *ll.Node) *ll.Node {
	var newList *ll.Node
	temp := lh
	temp2 := rh
	for temp != nil && temp2 != nil {
		if temp.Val < temp2.Val {
			newList = ll.AddNode(newList, temp.Val)
			temp = temp.Next
		} else {
			newList = ll.AddNode(newList, temp2.Val)
			temp2 = temp2.Next
		}
	}
	if temp == nil {
		for temp2 != nil {
			newList = ll.AddNode(newList, temp2.Val)
			temp2 = temp2.Next
		}
	}
	if temp2 == nil {
		for temp != nil {
			newList = ll.AddNode(newList, temp.Val)
			temp = temp.Next
		}
	}
	return newList
}

/*
func Run() {
	arr := []int{7, 6, 5, 4, 3, 2, 1}
	var head *ll.Node
	for i := 0; i < len(arr); i++ {
		head = ll.AddNode(head, arr[i])
	}
	ll.Traverse(head)
	sortedList := MergeSort(head)
	ll.Traverse(sortedList)
}*/
