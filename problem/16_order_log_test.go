package problem

import (
	"testing"
)

var orderLogTests = []struct {
	ids []int
	N   int
	i   int
	out int
}{
	{[]int{1, 2, 3, 4, 5}, 5, 5, 1},
	{[]int{1, 2, 3, 4, 5}, 5, 4, 2},
	{[]int{1, 2, 3, 4, 5}, 5, 3, 3},
	{[]int{1, 2, 3, 4, 5}, 5, 2, 4},
	{[]int{1, 2, 3, 4, 5}, 5, 1, 5},
	{[]int{1, 2, 3, 4, 5}, 4, 1, 5},
	{[]int{1, 2, 3, 4, 5}, 3, 1, 5},
	{[]int{1, 2, 3, 4, 5}, 2, 1, 5},
}

func TestOrderLog(t *testing.T) {
	for _, tc := range orderLogTests {
		ol := NewOrderLog(tc.N)
		for _, id := range tc.ids {
			ol.Record(id)
		}

		if actual, expected := ol.GetLast(tc.i), tc.out; actual != expected {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	}
}
