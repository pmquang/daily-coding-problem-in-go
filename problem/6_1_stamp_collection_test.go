package problem

import (
	"testing"

	"github.com/khoi/daily-coding-problem-in-go/helper"
)

var stampTests = []struct {
	in  []int
	out []int
}{
	{[]int{1, 1, 1, 3}, []int{1}},
	{[]int{1, 7, 3, 1, 7, 4, 5, 1, 7, 1}, []int{1, 7, 1}},
	{[]int{1, 2, 3, 4, 4}, []int{}},
	{[]int{5, 4, 4, 3, 3, 3}, []int{3}},
	{[]int{7, 7, 7, 7, 7}, []int{7, 7, 7}},
}

func TestStamp(t *testing.T) {
	for _, tc := range stampTests {
		if actual := Stamp(tc.in); !helper.Equal(actual, tc.out) {
			t.Errorf("expect %v got %v", tc.out, actual)
		}
	}
}
