package daily_coding_problem_in_go

import "testing"

var encodeTests = []struct {
	in  string
	out int
}{
	{"11", 2},
	{"111", 3},
	{"12", 2},
	{"123", 3},
}

func TestPossibleEncodedSolutions(t *testing.T) {
	for _, tc := range encodeTests {
		if actual := PossibleEncodedSolutions(tc.in); actual != tc.out {
			t.Errorf("%v expected %v, got %v", tc.in, tc.out, actual)
		}
	}
}
