package main

func (t *BinaryTree)second_nd_largest_item_bst(prev *BinaryTree) {
	if t.right != nil {
		t.right.second_nd_largest_item_bst(t)
	} else {
		for prev.right != nil {
			prev = prev.right
		}
	}
}
