package daily_coding_problem_in_go

import "time"

func Schedule(f func(), n int) {
	go func() {
		time.Sleep(time.Duration(n) * time.Millisecond)
		f()
	}()
}
