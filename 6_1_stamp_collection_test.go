package daily_coding_problem_in_go

import "testing"

var stampTests = []struct {
	in  []int
	out []int
}{
	{[]int{1, 7, 3, 1, 7, 4, 5, 1, 7, 1}, []int{1, 1, 7}},
	{[]int{1, 2, 3, 4, 4}, []int{}},
	{[]int{5, 4, 4, 3, 3, 3}, []int{3}},
	{[]int{7, 7, 7, 7, 7}, []int{7, 7, 7}},
}

func TestStamp(t *testing.T) {
	for _, tc := range stampTests {
		if Equal(Stamp(tc.in), tc.out) {
			t.Errorf("")
		}
	}
}
