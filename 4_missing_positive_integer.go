package daily_coding_problem_in_go

func MissingPositiveInt(arr []int) int {
	for i, e := range arr {
		if e <= 0 || e >= len(arr) {
			continue
		}
		arr[e-1], arr[i] = arr[i], arr[e-1]
	}

	for i := range arr {
		if arr[i] != i+1 {
			return i + 1
		}
	}

	return len(arr) + 1
}
