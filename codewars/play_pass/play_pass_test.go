package main

import "testing"

func TestPlayPass(t *testing.T) {
	pass := PlayPass("loool", 2)
	expectedPassphrase := "nqqqn"
	if pass != expectedPassphrase {
		t.Errorf("Passphrase was incorrect, got: %s, want: %s.", pass, expectedPassphrase)
	}

}
