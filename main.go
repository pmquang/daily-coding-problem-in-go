package main

import (
	"fmt"

	"github.com/khoi/daily-coding-problem-in-go/problem"
)

func main() {
	pi := problem.EstimatePiWithMonteCarlo(10000000)
	fmt.Println(pi)
}
