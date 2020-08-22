package tree

import (
	"fmt"
)

type node struct {
	val   int
	left  *node
	right *node
}

func newNode(val int) *node {
	return &node{
		val:   val,
		left:  nil,
		right: nil,
	}
}

func (root *node) Height(n *node, h int) int {
	if n == nil {
		return h
	}
	lh := n.Height(n.left, h+1)
	rh := n.Height(n.right, h+1)
	if lh > rh {
		return lh
	}
	return rh
}

func (root *node) PrintGivenLevel(n *node, expectedLvl, curLvl int) {
	if n == nil {
		return
	}
	if expectedLvl == curLvl {
		fmt.Printf(" %d", n.val)
	}
	n.PrintGivenLevel(n.left, expectedLvl, curLvl+1)
	n.PrintGivenLevel(n.right, expectedLvl, curLvl+1)
}

func (root *node) AddChild(val int) error {
	ch := &node{val: val}
	temp := root
	var par *node
	for temp != nil {
		par = temp
		if val < temp.val {
			temp = temp.left
		} else {
			temp = temp.right
		}
	}
	if val < par.val {
		par.left = ch
	} else {
		par.right = ch
	}
	return nil
}

func (root *node) leftView(n *node) {
	if n == nil {
		return
	}
	fmt.Println(n.val)
	root.leftView(n.left)
}

func (root *node) TraverseTree(n *node) error {
	if n == nil {
		return nil
	}
	fmt.Println(n.val)
	root.TraverseTree(n.left)
	root.TraverseTree(n.right)
	return nil
}

func (root *node) TraverseLevelOrderTree(index int, myList *[]*node) {
	if index == len(*myList) {
		return
	}
	par := (*myList)[index]
	if par.left != nil {
		*myList = append(*myList, par.left)
	}
	if par.right != nil {
		*myList = append(*myList, par.right)
	}
	root.TraverseLevelOrderTree(index+1, myList)
}

func (root *node) minimumSumRootToLeaf(n *node) int {
	if n.left == nil && n.right == nil {
		return n.val
	}
	// check if leaf
	var ls, rs int
	if n.left != nil {
		ls = n.val + root.minimumSumRootToLeaf(n.left)
	}
	if n.right != nil {
		rs = n.val + root.minimumSumRootToLeaf(n.right)
	}

	if rs == 0 {
		return ls
	}
	if ls == 0 {
		return ls
	}
	if ls <= rs {
		return ls
	} else {
		return rs
	}
}

func (root *node) sumTree(n *node, isSumtree *bool) {
	if n == nil {
		return
	}
	var sum int
	if n.left != nil && n.right != nil {
		sum = n.right.val + n.left.val
	} else if n.left == nil && n.right == nil {
		sum = n.val
	} else if n.right != nil {
		sum = n.right.val
	} else if n.left != nil {
		sum = n.left.val
	}

	if n.val != sum {
		*isSumtree = false
	}
	root.sumTree(n.left, isSumtree)
	if *isSumtree {
		root.sumTree(n.right, isSumtree)
	}
	return
}

func constructBasicBinaryTree() *node {
	root := newNode(100)
	root.left = newNode(50)
	root.right = newNode(50)
	root.left.left = newNode(10)
	root.left.right = newNode(41)
	// root.right.left = newNode(30)
	// root.right.right = newNode(20)
	return root
}

func Run() {
	//mylist := []int{20, 30, 5, 6, 25, 4, 15, 35, 5, 3, 4, 1, 5}
	// mylist := []int{4, 6}
	// root := newNode(10)
	// for _, i := range mylist {
	// 	root.AddChild(i)
	// }
	root := constructBasicBinaryTree()
	var isSumTree bool
	isSumTree = true
	root.sumTree(root, &isSumTree)
	fmt.Println(isSumTree)
	/*
		root.TraverseTree(root)
		mylvlorderList := []*node{root}
		root.TraverseLevelOrderTree(0, &mylvlorderList)
		mylvlorder := []int{}
		for _, nodes := range mylvlorderList {
			mylvlorder = append(mylvlorder, nodes.val)
		}
		fmt.Println(mylvlorder)
		height := root.Height(root, 0)
		fmt.Printf("Height of the tree is :%d\n", height)
		for i := 0; i < height; i++ {
			root.PrintGivenLevel(root, i, 0)
			fmt.Println()
		}
	*/

}

// 10 20 30 5 6 25
