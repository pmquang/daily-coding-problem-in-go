package daily_coding_problem_in_go

func ProductArray(arr []int) []int {
	product := arr[0]
	result := make([]int, len(arr))

	for _, e := range arr[1:] {
		product *= e
	}

	for i, e := range arr {
		result[i] = product / e
	}

	return result
}

func ProductArray2(arr []int) []int {
	size := len(arr)
	left := make([]int, size)
	right := make([]int, size)
	result := make([]int, size)

	for i := 0; i < size; i++ {
		left[i] = 1
		right[i] = 1
	}

	for i := 1; i < size; i++ {
		left[i] = arr[i-1] * left[i-1]
	}

	for i := size - 2; i >= 0; i-- {
		right[i] = arr[i+1] * right[i+1]
	}

	for i := 0; i < size; i++ {
		result[i] = left[i] * right[i]
	}

	return result
}
