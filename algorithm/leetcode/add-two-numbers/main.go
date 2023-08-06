package main

import "fmt"

func main() {
	l1 := &ListNode{
		3,
		&ListNode{
			7,
			nil,
		},
	}

	l2 := &ListNode{
		9,
		&ListNode{
			2,
			nil,
		},
	}

	res := addTwoNumbers(l1, l2)

	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1.Val == 0 && l1.Next == nil {
		return l2
	}

	if l2.Val == 0 && l2.Next == nil {
		return l1
	}

	currentResultNode := &ListNode{}
	headOfResultNode := currentResultNode
	addResult := 0
	tmp := int8(0)

	for l1 != nil && l2 != nil {

		addResult = l2.Val + l1.Val + currentResultNode.Val

		if addResult > 9 {
			tmp = 1
		}

		currentResultNode.Val += addResult % 10

		l1 = l1.Next
		l2 = l2.Next

		if l1 == nil && l2 == nil {

			if tmp == 1 {
				currentResultNode.Next = &ListNode{
					Val: int(tmp),
				}
			} else {
				currentResultNode.Next = nil
			}

			return headOfResultNode
		}

		currentResultNode.Next = &ListNode{
			Val: int(tmp),
		}
		currentResultNode = currentResultNode.Next
		tmp = 0
	}

	if l1 == nil {
		addOffset(l2, currentResultNode)

		return headOfResultNode
	}

	if l2 == nil {
		addOffset(l1, currentResultNode)
	}

	return headOfResultNode
}

func addOffset(l *ListNode, currentResultNode *ListNode) {
	tmp := int8(0)

	for l != nil {
		addResult := currentResultNode.Val + l.Val

		if addResult > 9 {
			tmp = 1
		}

		currentResultNode.Val = addResult % 10

		l = l.Next
		if l == nil {

			if tmp == 1 {
				currentResultNode.Next = &ListNode{
					Val: int(tmp),
				}
			} else {
				currentResultNode.Next = nil
			}

		}

		currentResultNode.Next = &ListNode{
			Val: int(tmp),
		}
		currentResultNode = currentResultNode.Next
		tmp = 0
	}
}
