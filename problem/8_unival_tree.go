package problem

import (
	"github.com/khoi/daily-coding-problem-in-go/helper"
)

func CountUnivalTree(tree *helper.BinaryTree) (int, bool) {
	if tree == nil {
		return 0, true
	}

	leftCount, isLeftUnival := CountUnivalTree(tree.Left)
	rightCount, isRightUnival := CountUnivalTree(tree.Right)
	total := leftCount + rightCount

	if isLeftUnival && isRightUnival {
		if tree.Left != nil && tree.Val != tree.Left.Val {
			return total, false
		}
		if tree.Right != nil && tree.Val != tree.Right.Val {
			return total, false
		}
		return total + 1, true
	}

	return total, false
}
