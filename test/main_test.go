package test

import "testing"

func TestAdd(t *testing.T) {
	a := 5
	b := 10

	exp := 15
	res := Sum(a, b)

	if exp != res {
		t.Errorf("actual and exp are not matching %d %d:", res, exp)
	}

}
