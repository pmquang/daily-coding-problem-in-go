package daily_coding_problem_in_go

import (
	"strconv"
)

func PossibleEncodedSolutions(message string) int {
	length := len(message)
	num, _ := strconv.Atoi(message)

	if length == 1 {
		return 1
	}

	if length == 2 && num > 0 && num < 27 {
		return 2
	}

	if length == 2 {
		return 1
	}

	return PossibleEncodedSolutions(message[1:]) + PossibleEncodedSolutions(message[2:])
}
