package problem

import (
	"testing"

	"github.com/khoi/daily-coding-problem-in-go/helper"
)

func TestCountUnivalTree(t *testing.T) {
	tree := &helper.BinaryTree{
		Val:  0,
		Left: &helper.BinaryTree{Val: 1},
		Right: &helper.BinaryTree{
			Val:   0,
			Left:  &helper.BinaryTree{Val: 1, Left: &helper.BinaryTree{Val: 1}, Right: &helper.BinaryTree{Val: 1}},
			Right: &helper.BinaryTree{Val: 0},
		},
	}
	expected := 5
	if actual, _ := CountUnivalTree(tree); actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
