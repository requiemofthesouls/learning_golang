package main

import (
	"fmt"
	"testing"
)

func TestEvaporator(t *testing.T) {
	answ := Evaporator(10, 10, 10)
	if answ != 22 {
		fmt.Println("Wrong answer! expected 22, got", answ)
	}
}
