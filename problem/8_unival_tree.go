//This problem was asked by Google.
//
//A unival tree (which stands for "universal value") is a tree where all nodes under it have the same value.
//
//Given the root to a binary tree, count the number of unival subtrees.
//
//For example, the following tree has 5 unival subtrees:
//
//   0
//  / \
// 1   0
//    / \
//   1   0
//  / \
// 1   1

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
