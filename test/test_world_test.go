package test

import "testing"

func TestFib(t *testing.T) {
	multi(t)
}

func simple(t *testing.T) {
	var (
		in     = 7
		expect = 14
	)
	actual := Fib(in)
	if expect != actual {
		t.Errorf("fib(%d),actual=%d,expect=%d", in, actual, expect)
	}
}

func multi(t *testing.T) {
	var fibTest = []struct {
		in     int
		expect int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
	}

	for _, val := range fibTest {
		actual := Fib(val.in)
		if actual != val.expect {
			t.Errorf("fib(%d),actual=%d,expect=%d", val.in, actual, val.expect)
		}
	}
}
