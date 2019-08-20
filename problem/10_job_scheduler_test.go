package problem

import (
	"sync"
	"testing"

	"github.com/khoi/daily-coding-problem-in-go/helper"
)

func TestSchedule(t *testing.T) {
	var results []int
	var expected []int
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)

		work := func(index int) func() {
			return func() {
				results = append(results, index)
				wg.Done()
			}
		}

		Schedule(work(i), i*100)
		expected = append(expected, i)
	}
	wg.Wait()

	if !helper.Equal(expected, results) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}
