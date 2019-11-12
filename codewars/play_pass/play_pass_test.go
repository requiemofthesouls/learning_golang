package main

import "testing"

func TestPlayPass(t *testing.T) {
	pass := PlayPass("loool", 2)
	if pass != "loool" {
		t.Errorf("Passphrase was incorrect, got: %s, want: %s.", pass, "zaazazaza")
	}

}
