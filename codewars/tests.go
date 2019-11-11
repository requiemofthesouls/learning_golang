package main

import "testing"

func TestCountBits(t *testing.T) {
	bin := CountBits(1234)
	if bin != 10011010010 {
		t.Errorf("BIN was incorrect, got: %d, want: %d.", bin, 10011010010)
	}
}
