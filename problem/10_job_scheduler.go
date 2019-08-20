package problem

import "time"

func Schedule(f func(), n int) {
	go func() {
		time.Sleep(time.Duration(n) * time.Millisecond)
		f()
	}()
}
