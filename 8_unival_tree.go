package daily_coding_problem_in_go

func CountUnivalTree(tree *BinaryTree) (int, bool) {
	if tree == nil {
		return 0, true
	}

	leftCount, isLeftUnival := CountUnivalTree(tree.left)
	rightCount, isRightUnival := CountUnivalTree(tree.right)
	total := leftCount + rightCount

	if isLeftUnival && isRightUnival {
		if tree.left != nil && tree.val != tree.left.val {
			return total, false
		}
		if tree.right != nil && tree.val != tree.right.val {
			return total, false
		}
		return total + 1, true
	}

	return total, false
}
