package problem

import (
	"testing"

	"github.com/khoi/daily-coding-problem-in-go/helper"
)

var xorLinkedListTests = [][]int{
	{1, 2, 3, 4, 5},
	{5, 4, 3, 2, 1},
}

func TestXorLinkedList(t *testing.T) {
	for _, tc := range xorLinkedListTests {
		a := NewXorLinkedList(tc[0])
		for _, e := range tc[1:] {
			a = a.Insert(e)
		}
		if actual, expected := a.ToSlice(), helper.Reverse(tc); !helper.Equal(actual, expected) {
			t.Errorf("expecting %v, got %v", expected, actual)
		}
	}
}
