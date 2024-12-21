package main

import (
	"slices"
	"testing"

	"aoc24/lib"
)

func TestParseLine(t *testing.T) {
	const in = "123: 5 4 66 77"
	expectRes := 123
	expectVals := []int{5, 4, 66, 77}
	r, v := parseLine(in)
	if r != expectRes {
		t.Errorf("expected result value to be %d, got %d", expectRes, r)
	}

	if !slices.Equal(expectVals, v) {
		t.Errorf("expected to get %v as remainder, got %v", expectVals, v)
	}
}

func TestValidTotal(t *testing.T) {
	const (
		in = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`
		expect       = 3749
		expectConcat = 11387
	)

	if res := validTotal(slices.Collect(lib.StringLines(in))); res != expect {
		t.Errorf("expected %d got %d", expect, res)
	}

	if res := validConcatTotal(slices.Collect(lib.StringLines(in))); res != expectConcat {
		t.Errorf("expected %d got %d", expectConcat, res)
	}
}
