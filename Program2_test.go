package main

import "testing"

func TestEx001(t *testing.T) {
	expect := 14
	result := RomantoArab("XIV")
	if result != expect {
		t.Error("Test01 is failed")
	}
}
