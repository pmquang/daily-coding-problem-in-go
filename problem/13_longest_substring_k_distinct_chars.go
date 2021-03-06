//This problem was asked by Amazon.
//
//Given an integer k and a string s, find the length of the longest substring that contains at most k distinct characters.
//
//For example, given s = "abcba" and k = 2, the longest substring with k distinct characters is "bcb".
package problem

import "github.com/khoi/daily-coding-problem-in-go/helper"

func CountLongestSubStrWithKUniqueChars(str string, k int) int {
	if len(str) == 0 {
		return 0
	}

	i := 0
	maxLength := 0
	dict := make(map[rune]int)

	for j, c := range str {
		val, ok := dict[c]

		if ok {
			dict[c] = val + 1
		} else {
			dict[c] = 1
		}

		if len(dict) <= k {
			maxLength = helper.Max(maxLength, j-i+1)
		} else {
			s := rune(str[i])
			dict[s] -= 1
			if dict[s] == 0 {
				delete(dict, s)
			}
			i++
		}
	}

	return maxLength
}
