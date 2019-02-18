package daily_coding_problem_in_go

import "testing"

var productArrayTests = []struct {
	in  []int
	out []int
}{
	{[]int{1, 2, 3, 4, 5}, []int{120, 60, 40, 30, 24}},
	{[]int{3, 2, 1}, []int{2, 3, 6}},
}

func TestProductArray(t *testing.T) {
	for _, tc := range productArrayTests {
		if actual := ProductArray(tc.in); !Equal(actual, tc.out) {
			t.Errorf("%v, expected %v, got %v", tc.in, tc.out, actual)
		}
	}
}
