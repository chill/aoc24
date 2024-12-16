package main

import (
	"slices"
	"testing"

	"aoc24/lib"
)

func TestParseInput(t *testing.T) {
	in := `47|53

61,13,29`

	input := parseInput(lib.StringLines(in))

	if len(input.rules) != 1 {
		t.Errorf("expected 1 rule but got %d", len(input.rules))
	} else if _, ok := input.rules[47][53]; !ok {
		t.Errorf("unexpected rule")
	}

	if len(input.updates) != 1 {
		t.Errorf("expected 1 update but got %d", len(input.updates))
	} else if !slices.Equal(input.updates[0], []int{61, 13, 29}) {
		t.Errorf("unexpected update")
	}
}
