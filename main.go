package main

import (
	"fmt"

	"github.com/khoi/daily-coding-problem-in-go/problem"
)

func main() {
	t := problem.NewTrie()
	t.Insert("deer")
	t.Insert("de")
	t.Insert("deal")
	b := t.Search("de")
	fmt.Println(b)
}
