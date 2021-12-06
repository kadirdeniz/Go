package main

import "testing"

type substractionTest struct {
	x, y, expected int
}

var substractionArr = []substractionTest{
	substractionTest{1, 2, 1},
	substractionTest{5, 10, 5},
	substractionTest{7, 8, 1},
}

func TestSubstraction(t *testing.T) {
	for _, test := range substractionArr {
		if response := substraction(test.x, test.y); response != test.expected {
			t.Errorf("Output %d not equal to expected %d", response, test.expected)
		}
	}
}
