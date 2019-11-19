package main

import "testing"

func TestDigPow(t *testing.T) {
	answ := DigPow(89, 1)
	expected := 1

	if answ != expected {
		t.Errorf("expected '%d' but got '%d'", expected, answ)
	}

	answ = DigPow(92, 1)
	expected = -1

	if answ != expected {
		t.Errorf("expected '%d' but got '%d'", expected, answ)
	}

	answ = DigPow(695, 2)
	expected = 2

	if answ != expected {
		t.Errorf("expected '%d' but got '%d'", expected, answ)
	}

	answ = DigPow(46288, 3)
	expected = 51

	if answ != expected {
		t.Errorf("expected '%d' but got '%d'", expected, answ)
	}

}
