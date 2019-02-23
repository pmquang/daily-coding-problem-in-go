package daily_coding_problem_in_go

import (
	"unsafe"
)

type XorNode struct {
	Val int
	Npx uintptr
}

// NewXorLinkedList returns a root node
func NewXorLinkedList(val int) *XorNode {
	return &XorNode{val, 0}
}

// Insert insert a value and return pointer to the new head node
func (head *XorNode) Insert(val int) *XorNode {
	newHeadNode := &XorNode{Val: val}
	ptr := unsafe.Pointer(head)

	newHeadNode.Npx = 0 ^ uintptr(ptr)
	newNodePtr := unsafe.Pointer(newHeadNode)

	head.Npx = uintptr(newNodePtr) ^ uintptr(head.Npx)

	return newHeadNode
}

func (head *XorNode) ToSlice() []int {
	var result []int
	prev := uintptr(0)
	this := unsafe.Pointer(head)
	result = append(result, head.Val)
	for prev^head.Npx != 0 {
		prev, this = uintptr(this), unsafe.Pointer(prev^head.Npx)
		head = (*XorNode)(this)
		result = append(result, head.Val)
	}
	return result
}
