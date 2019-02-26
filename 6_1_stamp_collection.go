package daily_coding_problem_in_go

func Stamp(arr []int) []int {
	m := make(map[int]int)
	result := make([]int, len(arr))

	for _, e := range arr {
		if _, ok := m[e]; ok {
			m[e]++
			continue
		}
		m[e] = 1
	}

	for k, v := range m {
		if v > 2 {
			for i := 0; i < v-2; i++ {
				result = append(result, k)
			}
		}
	}

	return result

}
