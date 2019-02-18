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