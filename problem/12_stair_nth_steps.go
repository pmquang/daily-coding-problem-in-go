package problem

func CountStairSteps(n int, steps []int) int {
	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	count := 0

	for _, step := range steps {
		count += CountStairSteps(n-step, steps)
	}

	return count
}
