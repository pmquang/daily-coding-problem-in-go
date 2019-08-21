//This problem was asked by Apple.
//
//Implement a job scheduler which takes in a function f and an integer n, and calls f after n milliseconds.

package problem

import "time"

func Schedule(f func(), n int) {
	go func() {
		time.Sleep(time.Duration(n) * time.Millisecond)
		f()
	}()
}
