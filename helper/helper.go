package helper

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func Reverse(s []int) []int {
	res := make([]int, len(s))
	copy(res, s)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
