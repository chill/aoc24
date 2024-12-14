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
	nums := lib.Ints(line)
	if diffSafe(diffs(nums)) {
		return true
	}

	// check all lists with one element knocked out
	for i := 0; i < len(nums); i++ {
		check := lib.SafeDel(nums, i)
		if diffSafe(diffs(check)) {
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

func diffs(nums []int) []int {
	prev := nums[0]
	diffs := make([]int, 0, 2)
	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		diff := curr - prev
		prev = curr
		diffs = append(diffs, diff)
	}

	return diffs
}
