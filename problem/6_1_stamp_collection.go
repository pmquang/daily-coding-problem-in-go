package problem

func Stamp(arr []int) []int {
	m := make(map[int]int)
	var result []int

	for _, e := range arr {
		if _, ok := m[e]; !ok {
			m[e] = 1
			continue
		}

		if m[e] == 2 {
			result = append(result, e)
			continue
		}

		m[e]++
	}

	return result

}
