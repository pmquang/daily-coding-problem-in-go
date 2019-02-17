package daily_coding_problem_in_go

import "testing"

var tests = []struct {
	arr      []int
	k        int
	expected bool
}{
	{[]int{10, 15, 3, 7}, 17, true},
	//{[]int{10, 15, 3, 7}, 10, true},
	//{[]int{10, 15, 3, 7}, 3, false},
	//{[]int{10, 15, 3, 7}, 26, false},
}

func TestTwoSum(t *testing.T) {
	for _, tc := range tests {
		if actual := TwoSum(tc.arr, tc.k); actual != tc.expected {
			t.Errorf("arr = %v, k = %v, got %v, want %v", tc.arr, tc.k, actual, tc.expected)
		}
	}
}
