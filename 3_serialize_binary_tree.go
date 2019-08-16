package daily_coding_problem_in_go

import "encoding/json"

func Serialize(n *Node) string {
	b, _ := json.Marshal(n)
	return string(b)
}

func Deserialize(data string) *Node {
	n := new(Node)
	_ = json.Unmarshal([]byte(data), &n)
	return n
}
