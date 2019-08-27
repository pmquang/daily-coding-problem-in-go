// This problem was asked by Twitter.
//
// You run an e-commerce website and want to record the last N order ids in a log.
// Implement a data structure to accomplish this, with the following API:
//
// record(order_id): adds the order_id to the log
// get_last(i): gets the ith last element from the log. i is guaranteed to be smaller than or equal to N.
// You should be as efficient with time and space as possible.

package problem

type OrderLog struct {
	orders  []int
	currIdx int
	N       int
}

func NewOrderLog(n int) *OrderLog {
	return &OrderLog{
		orders:  make([]int, n),
		currIdx: -1,
		N:       n,
	}
}

func (ol *OrderLog) Record(id int) {
	ol.currIdx++
	if ol.currIdx+1 > ol.N {
		ol.currIdx = 0
	}
	ol.orders[ol.currIdx] = id
}

func (ol *OrderLog) GetLast(i int) int {
	idx := ol.currIdx - i + 1
	if idx < 0 {
		idx += ol.N
	}

	return ol.orders[idx]
}
