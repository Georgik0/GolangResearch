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

func (t *stack_tree)push(elem *BinaryTree) {
	*t = append(*t, *elem)
}

func (t *stack_tree)back() {
	*t = (*t)[:(len(*t) - 1)]
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
}
