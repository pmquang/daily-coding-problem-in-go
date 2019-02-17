package daily_coding_problem_in_go

func TwoSum(nums []int, k int) bool {
	m := make(map[int]struct{})

	for _, e := range nums {
		if _, ok := m[k-e]; ok {
			return true
		}
		m[e] = struct{}{}
	}

	return false
}
