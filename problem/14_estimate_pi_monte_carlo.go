package problem

import (
	"math/rand"
	"time"
)

func EstimatePiWithMonteCarlo(iterations int) float64 {
	rand.Seed(time.Now().UnixNano())
	inner := 0
	for i := 0; i < iterations; i++ {
		x := rand.Float64()
		y := rand.Float64()
		if x*x+y*y <= 1 {
			inner += 1
		}
	}
	return 4 * float64(inner) / float64(iterations)
}
