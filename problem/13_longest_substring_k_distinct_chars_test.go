package problem

import (
	"testing"
)

var longestSubstrWithKUniqueCharsTests = []struct {
	s   string
	k   int
	out int
}{
	{"abcba", 2, 3},
	{"aabbcc", 1, 2},
	{"aabbcc", 2, 4},
	{"aabbcc", 3, 6},
}

func TestCountLongestSubStrWithKUniqueChars(t *testing.T) {
	for _, tc := range longestSubstrWithKUniqueCharsTests {
		if actual, expected := CountLongestSubStrWithKUniqueChars(tc.s, tc.k), tc.out; actual != expected {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	}
}
