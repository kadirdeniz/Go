package main

import "testing"

type additionTest struct {
	x, y, expected int
}

var additionArr = []additionTest{
	additionTest{1, 2, 3},
	additionTest{5, 10, 15},
	additionTest{7, 8, 15},
}

func TestAdding(t *testing.T) {
	for _, test := range additionArr {
		if response := addition(test.x, test.y); response != test.expected {
			t.Errorf("Output %d not equal to expected %d", response, test.expected)
		}
	}
}
