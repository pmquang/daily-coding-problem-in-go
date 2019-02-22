package daily_coding_problem_in_go

import "testing"

func TestCar(t *testing.T) {
	if actual := Car(Cons(3,4)); actual != 3 {
		t.Errorf("Expected 3 - got %v", actual)
	}
}

func TestCdr(t *testing.T) {
	if actual := Cdr(Cons(3,4)); actual != 4 {
		t.Errorf("Expected 4 - got %v", actual)
	}
}