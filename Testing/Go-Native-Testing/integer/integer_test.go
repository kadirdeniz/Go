package integer

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
func TestMinus(t *testing.T) {
	sum := Minus(2, 1)
	expected := 1

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
