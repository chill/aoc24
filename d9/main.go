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
	m2 := slices.Clone(m)
	fmt.Println(checksum(compact(m)))
	defragged := compactDefrag(m2)
	fmt.Println(checksum(defragged))
}

func loadMap(in iter.Seq[rune]) []int {
	runes := slices.Collect(in)

	res := make([]int, 0, len(runes)/2)
	for i, r := range runes {
		v := lib.Atoi(string(r))
		if i%2 == 0 { // file
			res = append(res, lib.RepeatSlice((i/2)+1, v)...)
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

func compactDefrag(o []int) []int {
	for i := len(o) - 1; i >= 0; i-- { // right to left, val by val
		if o[i] == 0 {
			continue
		}

		file := getFile(i, o)
		i -= len(file) - 1 // this puts i at the beginning of the file, next iteration we'll move off it

		for j := 0; j < i; j++ { // left to right, find the empty slots, don't go further than file pos
			if o[j] > 0 {
				continue
			}

			runLen := runLength(j, o)
			if len(file) > runLen {
				j += runLen - 1
				continue
			}

			// found a slot we can fit this file in, let's do it
			for k := j; k < j+len(file); k++ {
				o[k] = file[0]
			}

			for k := i; k < i+len(file); k++ {
				o[k] = 0
			}
			break
		}
	}

	return o
}

func getFile(i int, o []int) []int {
	count := 1
	for j := i - 1; j >= 0; j-- {
		if o[j] != o[i] {
			break
		}

		count++
	}

	return lib.RepeatSlice(o[i], count)
}

func runLength(i int, o []int) int {
	count := 1
	for j := i + 1; j < len(o); j++ {
		if o[j] != o[i] {
			break
		}

		count++
	}

	return count
}

func checksum(o []int) uint64 {
	var sum uint64
	for i, v := range o {
		if v == 0 {
			// empty slot, done because we assume we're compacted
			continue
		}

		sum += uint64(i * (v - 1)) // minus one because we represent empty as 0, so the actual values are offset by +1
	}

	return sum
}
