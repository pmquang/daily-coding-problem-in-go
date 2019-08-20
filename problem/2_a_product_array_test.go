package problem

import (
	"testing"

	"github.com/khoi/daily-coding-problem-in-go/helper"
)

var productArrayTests = []struct {
	in  []int
	out []int
}{
	{[]int{1, 2, 3, 4, 5}, []int{120, 60, 40, 30, 24}},
	{[]int{3, 2, 1}, []int{2, 3, 6}},
}

func TestProductArray(t *testing.T) {
	for _, tc := range productArrayTests {
		if actual := ProductArray(tc.in); !helper.Equal(actual, tc.out) {
			t.Errorf("%v, expected %v, got %v", tc.in, tc.out, actual)
		}
	}
}

func TestProductArray2(t *testing.T) {
	for _, tc := range productArrayTests {
		if actual := ProductArray2(tc.in); !helper.Equal(actual, tc.out) {
			t.Errorf("%v, expected %v, got %v", tc.in, tc.out, actual)
		}
	}
}
