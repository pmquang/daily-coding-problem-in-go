package daily_coding_problem_in_go

import "testing"

var largestNonAdjacentSumTests = []struct {
	in  []int
	out int
}{
	{[]int{2, 4, 6, 2, 5}, 13},
	{[]int{5, 5, 10, 40, 50, 35}, 80},
	{[]int{5, 1, 1, 5}, 10},
	{[]int{5, 5, 10, 100, 10, 5}, 110},
	{[]int{1, 2, 3}, 4},
	{[]int{1, 20, 3}, 20},
}

func TestLargestNonAdjacentSum(t *testing.T) {
	for _, tc := range largestNonAdjacentSumTests {
		if actual := LargestNonAdjacentSum(tc.in); actual != tc.out {
			t.Errorf("%v expected %v, got %v", tc.in, tc.out, actual)
		}
	}
}
