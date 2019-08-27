package problem

import (
	"testing"

	"github.com/khoi/daily-coding-problem-in-go/helper"
)

var orderLogTests = []struct {
	ids []int
	i   int
	out []int
}{
	{[]int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
	{[]int{1, 2, 3, 4, 5}, 4, []int{2, 3, 4, 5}},
	{[]int{1, 2, 3, 4, 5}, 3, []int{3, 4, 5}},
	{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5}},
	{[]int{1, 2, 3, 4, 5}, 1, []int{5}},
	{[]int{1, 2, 3, 4, 5}, 0, []int{}},
}

func TestOrderLog(t *testing.T) {
	for _, tc := range orderLogTests {
		ol := NewOrderLog(10)
		for _, id := range tc.ids {
			ol.Record(id)
		}

		if actual, expected := ol.GetLast(tc.i), tc.out; !helper.Equal(actual, expected) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	}
}
