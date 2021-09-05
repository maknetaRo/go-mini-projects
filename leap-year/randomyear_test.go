package main

import "testing"

func TestIsLeapYear(t *testing.T) {
	got := IsLeapYear(2021)
	want := false

	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}

