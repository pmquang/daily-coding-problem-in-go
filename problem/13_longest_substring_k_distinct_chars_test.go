package problem

import (
	"testing"
)

var longestSubstrTests = []struct {
	s   string
	k   int
	out int
}{
	{"abcba", 2, 3},
	{"aabbcc", 1, 2},
	{"aabbcc", 2, 4},
	{"aabbcc", 3, 6},
	{"aaabbb", 3, 0},
}

func TestCountLongestSubStr(t *testing.T) {
	for _, tc := range longestSubstrTests {
		if actual, expected := CountLongestSubStr(tc.s, tc.k), tc.out; actual != expected {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	}
}
