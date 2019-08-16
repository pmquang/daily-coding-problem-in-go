package daily_coding_problem_in_go

import "testing"

func TestCountUnivalTree(t *testing.T) {
	tree := &BinaryTree{
		val:  0,
		left: &BinaryTree{val: 1},
		right: &BinaryTree{
			val:   0,
			left:  &BinaryTree{val: 1, left: &BinaryTree{val: 1}, right: &BinaryTree{val: 1}},
			right: &BinaryTree{val: 0},
		},
	}
	expected := 5
	if actual := CountUnivalTree(tree); actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
