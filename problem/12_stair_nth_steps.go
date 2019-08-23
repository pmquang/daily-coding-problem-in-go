package problem

func CountStairSteps(n int, steps []int) int {
	cache := make([]int, n+1)
	cache[0] = 1

	for i := 1; i < n; i++ {
		cache[i] = 0
	}

	for i := 1; i <= n; i++ {
		for _, step := range steps {
			if i-step >= 0 {
				cache[i] += cache[i-step]
			}
		}
	}

	return cache[n]
}
