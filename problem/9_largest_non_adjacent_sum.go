package problem

import "github.com/khoi/daily-coding-problem-in-go/helper"

func LargestNonAdjacentSum(l []int) int {
	if len(l) == 0 {
		return 0
	}
	maxIncludeLast, maxExcludeLast := l[0], 0
	for i := 1; i < len(l); i++ {
		inclTmp := maxIncludeLast
		maxIncludeLast = maxExcludeLast + l[i]
		maxExcludeLast = helper.Max(maxExcludeLast, inclTmp)
	}
	return helper.Max(maxIncludeLast, maxExcludeLast)
}
