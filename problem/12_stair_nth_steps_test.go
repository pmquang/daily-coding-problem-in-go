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
}

func TestCountStairSteps(t *testing.T) {
	for _, tc := range stairStepsTests {
		if actual, expected := CountStairSteps(tc.n, tc.possibleSteps), tc.out; actual != expected {
			t.Errorf("%v expected %v, got %v", tc.n, expected, actual)
		}
	}
}
