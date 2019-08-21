//This problem was asked by Airbnb.
//
//Given a list of integers, write a function that returns the largest sum of non-adjacent numbers. Numbers can be 0 or negative.
//
//For example, [2, 4, 6, 2, 5] should return 13, since we pick 2, 6, and 5. [5, 1, 1, 5] should return 10, since we pick 5 and 5.
//
//Follow-up: Can you do this in O(N) time and constant space?

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
