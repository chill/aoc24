package main

import "testing"

func TestIsSafe(t *testing.T) {
	tests := map[string]bool{
		"7 6 4 2 1":         true,
		"1 2 7 8 9":         false,
		"9 7 6 2 1":         false,
		"1 3 2 4 5":         true,
		"8 6 4 4 1":         true,
		"1 3 6 7 9":         true,
		"48 47 45 42 43 42": true,
		"48 47 45 43 42":    true, // knock out the element prior to first failure
		"48 47 45 42 43":    true, // knock out the last element
		"4 5 4 3 2":         true,
	}

	for test, expect := range tests {
		if result := lineSafe(test); result != expect {
			t.Errorf("expected input %s to have output %v but got %v", test, expect, result)
		}
	}
}
