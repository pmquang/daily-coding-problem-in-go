package daily_coding_problem_in_go

import "testing"

var missingPositiveIntTests = []struct {
	in  []int
	out int
}{
	{[]int{3, 4, -1, 1}, 2},
	{[]int{1, 2, 0}, 3},
	{[]int{1, 2, 3}, 4},
	{[]int{3, 2, 1}, 4},
	{[]int{-1, 4, 2, 3, 1}, 5},
	{[]int{3, 2, 4, -1, 1}, 5},
}

func TestMissingPositiveInt(t *testing.T) {
	for _, tc := range missingPositiveIntTests {
		dup := make([]int, len(tc.in)) // dupe it since we modify the input in-place
		copy(dup, tc.in)
		if actual := MissingPositiveInt(tc.in); actual != tc.out {
			t.Errorf("%v, expecting %v, got %v", dup, tc.out, actual)
		}
	}
}
