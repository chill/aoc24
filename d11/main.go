package main

import (
	"fmt"
	"maps"
	"slices"
	"strconv"

	"aoc24/lib"
)

func main() {
	line := slices.Collect(lib.ByLines("./input.txt"))[0]
	ints := lib.IntsWords(line)
	stones := map[int]int{}
	for _, i := range ints {
		stones[i]++
	}
	for range 25 {
		stones = blink(stones)
	}
	fmt.Println(tally(stones))
	for range 50 {
		stones = blink(stones)
	}
	fmt.Println(tally(stones))
}

func tally(stones map[int]int) int {
	var count int
	for c := range maps.Values(stones) {
		count += c
	}
	return count
}

func blink(stones map[int]int) map[int]int {
	newStones := map[int]int{}
	for stone, count := range stones {
		res := blinkStone(stone)
		for _, r := range res {
			newStones[r] += count
		}
	}
	return newStones
}

func blinkStone(v int) []int {
	if v == 0 {
		return []int{1}
	}

	s := strconv.Itoa(v)
	if len(s)%2 == 0 {
		left := lib.Atoi(s[:len(s)/2])
		right := lib.Atoi(s[len(s)/2:])
		return []int{left, right}
	}

	return []int{v * 2024}
}
