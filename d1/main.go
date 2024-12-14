package main

import (
	"fmt"
	"sort"

	"aoc24/lib"
)

func main() {
	lines := lib.ByLines("./input.txt")

	left := make([]int, 0, 1)
	right := make([]int, 0, 1)
	rightCounts := map[int]int{}

	for line := range lines {
		ints := lib.Ints(line)
		left = append(left, ints[0])
		right = append(right, ints[1])
		rightCounts[ints[1]]++
	}

	sort.Ints(left)
	sort.Ints(right)

	totalDistance := 0
	simScore := 0
	for i, v := range left {
		d := right[i] - v
		if d < 0 {
			d *= -1
		}

		totalDistance += d
		simScore += v * rightCounts[v]
	}

	fmt.Printf("total distance is %d\n", totalDistance)
	fmt.Printf("similarity score is %d\n", simScore)
}
