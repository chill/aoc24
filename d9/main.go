package main

import (
	"fmt"
	"iter"
	"slices"

	"aoc24/lib"
)

func main() {
	input := lib.ByRunes("./input.txt")
	m := loadMap(input)
	fmt.Println(checksum(compact(m)))
}

func loadMap(in iter.Seq[rune]) []int {
	fileID := 1
	runes := slices.Collect(in)

	res := make([]int, 0, len(runes)/2)
	for i, r := range runes {
		v := lib.Atoi(string(r))
		if i%2 == 0 { // file
			res = append(res, lib.RepeatSlice(fileID, v)...)
			fileID++
			continue
		}

		// empty
		res = append(res, make([]int, v)...) // append a slice of 0 vals of length v
	}

	return res
}

func compact(o []int) []int {
	for i := len(o) - 1; i >= 0; i-- { // right to left, val by val
		if o[i] == 0 {
			continue
		}

		moved := false
		for j := 0; j < i; j++ { // left to right, find the empty slots, don't go further than original pos
			if o[j] > 0 {
				continue
			}

			o[j] = o[i]
			o[i] = 0
			moved = true
		}

		if !moved {
			// no more free space, no point continuing
			break
		}
	}

	return o
}

func checksum(o []int) int {
	var sum int
	for i, v := range o {
		if v == 0 {
			// empty slot, done because we assume we're compacted
			break
		}

		sum += i * (v - 1) // minus one because we represent empty as 0, so the actual values are offset by +1
	}

	return sum
}
