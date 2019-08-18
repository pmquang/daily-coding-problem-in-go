package daily_coding_problem_in_go

func LargestNonAdjacentSum(l []int) int {
	if len(l) == 0 {
		return 0
	}
	incl, excl := l[0], 0
	for i := 1; i < len(l); i++ {
		inclTmp := incl
		incl = excl + l[i]
		excl = Max(excl, inclTmp)
	}
	return Max(incl, excl)
}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
