package problem

import "encoding/json"

type Node struct {
	Val   string `json:"val"`
	Left  *Node  `json:"left"`
	Right *Node  `json:"right"`
}

func Serialize(n *Node) string {
	b, _ := json.Marshal(n)
	return string(b)
}

func Deserialize(data string) *Node {
	n := new(Node)
	_ = json.Unmarshal([]byte(data), &n)
	return n
}
