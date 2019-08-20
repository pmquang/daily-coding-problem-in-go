package problem

import "testing"

func TestDeserialize(t *testing.T) {
	node := &Node{
		Val:   "root",
		Left:  &Node{Val: "left", Left: &Node{Val: "left.left"}},
		Right: &Node{Val: "right"},
	}

	newNode := Deserialize(Serialize(node))

	if newNode.Left.Left.Val != "left.left" {
		t.Errorf("expecting left.left")
	}
}
