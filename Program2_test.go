package main

import "testing"

func TestEx001(t *testing.T) {
	expect := 1990
	result := RomantoArab("MCMXC")
	if result != expect {
		t.Error("Test001 is failed")
	}
}

func TestEx002(t *testing.T) {
	expect := 2008
	result := RomantoArab("MMVIII")
	if result != expect {
		t.Error("Test002 is failed")
	}
}

func TestEx003(t *testing.T) {
	expect := 99
	result := RomantoArab("XCIX")
	if result != expect {
		t.Error("Test003 is failed")
	}
}

func TestEx004(t *testing.T) {
	expect := 47
	result := RomantoArab("XLVII")
	if result != expect {
		t.Error("Test004 is failed")
	}
}
