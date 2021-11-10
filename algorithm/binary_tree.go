package main

import (
	"errors"
	"fmt"
)

type BinaryTree struct {
	left *BinaryTree
	value int
	right *BinaryTree
	idx int
}

func (t *BinaryTree) Insert(v, idx int) error {
	if t == nil {
		return errors.New("Tree is nil")
	}
	if t.value == v {
		return errors.New("This value already exist")
	}
	if v < t.value {
		if t.left != nil {
			return t.left.Insert(v, idx)
		} else {
			t.left = &BinaryTree{value: v, idx: idx}
		}
	} else {
		if t.right != nil {
			return t.right.Insert(v, idx)
		} else {
			t.right = &BinaryTree{value: v, idx: idx}
		}
	}
	return nil
}

func (t *BinaryTree)detourTree() {
	if t == nil {
		return
	}
	t.left.detourTree()
	fmt.Printf("index: %v\tvalue: %v\n", t.idx, t.value)
	t.right.detourTree()
}

type stack_tree []BinaryTree

var h_max int = 0
var h_current int = 0
var is_balanced = true

func (t *stack_tree)push(elem *BinaryTree) {
	*t = append(*t, *elem)
}

func (t *stack_tree)back() {
	*t = (*t)[:(len(*t) - 1)]
}

func (t *BinaryTree)checkBalanced() {
	if t.left != nil {
		h_current++
		t.left.checkBalanced()
	}
	if t.right != nil {
		h_current++
		t.right.checkBalanced()
	} else {
		if h_max != 0 && h_current - h_max > 1 {
			is_balanced = false
		}
		if h_max == 0 {
			h_max = h_current
		}
		h_current--
		return
	}
}

func checkAll(t *BinaryTree, i int) {
	is_balanced = true
	t.checkBalanced()
	if is_balanced {
		fmt.Println(i, "balanced")
	} else {
		fmt.Println(i, "not balanced")
	}
}
//func (t *BinaryTree)checkBalanced(prevDeep int) bool {
//	if t.left == nil && t.right == nil {
//		return false
//	}
//	if t.left == nil && t.right != nil {
//
//	}
//}

func main() {
	node := &BinaryTree{value: 1}
	vector := []int{1,2,3,4,5,6,7,8,9}
	for idx, val := range vector {
		node.Insert(val, idx)
	}
	fmt.Println("Detour node")
	node.detourTree()
	t1 := &BinaryTree{nil, 1, nil, 1}

	t2 := &BinaryTree{nil, 1, nil, 1}
	t2.right = &BinaryTree{nil, 2, nil, 2}

	t3 := &BinaryTree{nil, 1, nil, 1}
	t3.right = &BinaryTree{nil, 2, nil, 2}
	t3.right.right = &BinaryTree{nil, 3, nil, 3}

	t4 := &BinaryTree{nil, 1, nil, 1}
	t4.left = &BinaryTree{nil, 2, nil, 2}
	t4.right = &BinaryTree{nil, 3, nil, 3}
	t4.right.right = &BinaryTree{nil, 4, nil, 4}

	t5 := &BinaryTree{nil, 1, nil, 1}
	t5.left = &BinaryTree{nil, 2, nil, 2}
	t5.right = &BinaryTree{nil, 3, nil, 3}
	t5.right.right = &BinaryTree{nil, 4, nil, 4}
	t5.right.right.right = &BinaryTree{nil, 5, nil, 5}


	t6 := &BinaryTree{nil, 1, nil, 1}
	t6.left = &BinaryTree{nil, 2, nil, 2}
	t6.left.left = &BinaryTree{nil, 4, nil,3}
	t6.left.right = &BinaryTree{nil, 5, nil, 4}
	t6.left.right.right = &BinaryTree{nil, 6, nil, 5}
	t6.right = &BinaryTree{nil, 3, nil, 6}
	t6.right.right = &BinaryTree{nil, 7, nil, 7}
	t6.right.right.left = &BinaryTree{nil, 8, nil, 8}
	t6.right.right.right = &BinaryTree{nil, 9, nil, 9}
	t6.right.right.right.right = &BinaryTree{nil, 10, nil, 10}
	fmt.Println("\nDetour t6")
	t6.detourTree()
	checkAll(t1, 1)
	checkAll(t2, 2)
	checkAll(t3, 3)
	checkAll(t4, 4)
	checkAll(t5, 5)
	checkAll(t6, 6)
}
