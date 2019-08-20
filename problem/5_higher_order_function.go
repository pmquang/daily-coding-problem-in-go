package problem

type Pair struct {
	First  int
	Second int
}

func Cons(a, b int) func() Pair {
	return func() Pair {
		return Pair{a, b}
	}
}

func Car(f func() Pair) int {
	return f().First
}

func Cdr(f func() Pair) int {
	return f().Second
}
