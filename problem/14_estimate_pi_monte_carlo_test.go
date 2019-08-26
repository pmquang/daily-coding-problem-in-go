package problem

import (
	"math"
	"testing"
)

var estimatePiWithMonteCarloTests = []struct {
	iterations int
}{
	{10000000},
	{100000000},
}

func TestEstimatePiWithMonteCarlo(t *testing.T) {
	for _, tc := range estimatePiWithMonteCarloTests {
		if actual := EstimatePiWithMonteCarlo(tc.iterations); !almostEqual(actual, math.Pi, 0.001) {
			t.Errorf("iteration %v got pi value of %v", tc.iterations, actual)
		}
	}
}

func BenchmarkEstimatePiWithMonteCarlo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range estimatePiWithMonteCarloTests {
			EstimatePiWithMonteCarlo(tc.iterations)
		}
	}
}

func almostEqual(a, b, delta float64) bool {
	return math.Abs(a-b) <= delta
}
