package daily_coding_problem_in_go

type Node struct {
	Val   string `json:"val"`
	Left  *Node  `json:"left"`
	Right *Node  `json:"right"`
}
