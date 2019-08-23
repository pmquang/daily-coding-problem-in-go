package problem

import "testing"

var stairStepsTests = []struct {
	n             int
	possibleSteps []int
	out           int
}{
	{1, []int{1, 2}, 1},
	{2, []int{1, 2}, 2},
	{3, []int{1, 2}, 3},
	{4, []int{1, 2}, 5},
	{5, []int{1, 2}, 8},
	{1, []int{1, 3, 5}, 1},
	{2, []int{1, 3, 5}, 1},
	{3, []int{1, 3, 5}, 2},
	{4, []int{1, 3, 5}, 3},
	{5, []int{1, 3, 5}, 5},
	{6, []int{1, 3, 5}, 8},
	{7, []int{1, 3, 5}, 12},
	{8, []int{1, 3, 5}, 19},
	{9, []int{1, 3, 5}, 30},
	{1, []int{1, 2, 3}, 1},
	{2, []int{1, 2, 3}, 2},
	{3, []int{1, 2, 3}, 4},
	{4, []int{1, 2, 3}, 7},
	{5, []int{1, 2, 3}, 13},
	{6, []int{1, 2, 3}, 24},
	{7, []int{1, 2, 3}, 44},
	{8, []int{1, 2, 3}, 81},
	{9, []int{1, 2, 3}, 149},
}

func TestCountStairSteps(t *testing.T) {
	for _, tc := range stairStepsTests {
		if actual, expected := CountStairSteps(tc.n, tc.possibleSteps), tc.out; actual != expected {
			t.Errorf("%v expected %v, got %v", tc.n, expected, actual)
		}
	}
}
