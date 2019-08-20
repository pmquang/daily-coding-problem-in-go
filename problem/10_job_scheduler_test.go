package problem

import (
	"sync"
	"testing"

	"github.com/khoi/daily-coding-problem-in-go/helper"
)

func TestSchedule(t *testing.T) {
	sliceMux := &sync.Mutex{}
	var results []int
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)

		work := func(index int) func() {
			return func() {
				sliceMux.Lock()
				results = append(results, index)
				sliceMux.Unlock()
				wg.Done()
			}
		}

		Schedule(work(i), i*100)
	}
	wg.Wait()

	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	if !helper.Equal(expected, results) {
		t.Errorf("expected %v, got %v", expected, results)
	}
}
