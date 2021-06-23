package main

import "testing"

func TestEx01(t *testing.T) {
	expect := "XIV"
	result := inttoroman(14)
	if result != expect {
		t.Error("Test01 is failed")
	}
}

func TestEx02(t *testing.T) {
	expect := "MCMXC"
	result := inttoroman(1990)
	if result != expect {
		t.Error("Test02 is failed")
	}
}

func TestEx03(t *testing.T) {
	expect := "MMVIII"
	result := inttoroman(2008)
	if result != expect {
		t.Error("Test03 is failed")
	}
}

func TestEx04(t *testing.T) {
	expect := "XCIX"
	result := inttoroman(99)
	if result != expect {
		t.Error("Test04 is failed")
	}
}

func TestEx05(t *testing.T) {
	expect := "XLVII"
	result := inttoroman(47)
	if result != expect {
		t.Error("Test05 is failed")
	}
}

func TestEx06(t *testing.T) {
	expect := "XLII"
	result := inttoroman(42)
	if result != expect {
		t.Error("Test06 is failed")
	}
}
