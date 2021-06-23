package main

import "testing"

func TestEx01(t *testing.T) {
	expect := "XIV"
	result := inttoroman(14)
	if result != expect {
		t.Error("Test01 is failed")
	}
}