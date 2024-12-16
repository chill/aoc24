package main

import (
	"fmt"

	"aoc24/lib"
)

func main() {
	lines := lib.ByLines("./input.txt")
	safe := 0
	for line := range lines {
		if lineSafe(line) {
			safe++
		}
	}

	fmt.Printf("%d reports are safe\n", safe)
}

func lineSafe(line string) bool {
	nums := lib.IntsWords(line)

	// check all lists with one element knocked out, starting with skip at -1 (no skip)
	for skip := -1; skip < len(nums); skip++ {
		if diffSafe(diffs(nums, skip)) {
			return true
		}
	}

	return false

}

func diffSafe(diffs []int) bool {
	sign := diffs[0] > 0
	for _, v := range diffs {
		if lib.Abs(v) > 3 || v == 0 {
			return false
		}

		currSign := v > 0
		if currSign != sign {
			return false
		}
	}

	return true
}

// skip below 0 has no effect in this logic
func diffs(nums []int, skip int) []int {
	start := 0
	if skip == 0 && len(nums) == 1 {
		panic("cannot skip index 0 in 1-length slice")
	} else if skip == 0 {
		start = 1
	}

	prev := nums[start]
	diffs := make([]int, 0, 2)
	for i := start + 1; i < len(nums); i++ {
		if i == skip {
			continue
		}

		curr := nums[i]
		diff := curr - prev
		prev = curr
		diffs = append(diffs, diff)
	}

	return diffs
}
