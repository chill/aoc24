package main

import (
	"fmt"
	"iter"
	"sort"
	"strings"

	"aoc24/lib"
)

func main() {
	lines := lib.ByLines("./input.txt")
	input := parseInput(lines)

	var total int
	var totalFixed int
	for _, update := range input.updates {
		toSort := &sortUpdate{
			nums:  update,
			rules: input.rules,
		}

		if !sort.IsSorted(toSort) {
			sort.Sort(toSort)
			mid := toSort.nums[len(toSort.nums)/2]
			totalFixed += mid
			continue
		}

		mid := toSort.nums[len(toSort.nums)/2]
		total += mid
	}
	fmt.Printf("part one: %d\npart two: %d\n", total, totalFixed)
}

type sortUpdate struct {
	nums  []int
	rules map[int]map[int]struct{}
}

func (u *sortUpdate) Len() int {
	return len(u.nums)
}

func (u *sortUpdate) Less(i, j int) bool {
	a, b := u.nums[i], u.nums[j]
	_, ok := u.rules[a][b] // if j appears in the rules for i, i should come before j
	return ok
}

func (u *sortUpdate) Swap(i, j int) {
	u.nums[i], u.nums[j] = u.nums[j], u.nums[i]
}

type input struct {
	rules   map[int]map[int]struct{}
	updates [][]int
}

func parseInput(lines iter.Seq[string]) *input {
	var upd bool
	rules := map[int]map[int]struct{}{}
	updates := make([][]int, 0)
	for line := range lines {
		if strings.TrimSpace(line) == "" {
			upd = true
			continue
		}

		if upd {
			updates = append(updates, parseLine(line, ","))
		} else {
			r := parseLine(line, "|")
			key := rules[r[0]]
			if key == nil {
				key = map[int]struct{}{}
				rules[r[0]] = key
			}
			key[r[1]] = struct{}{}
		}
	}

	return &input{
		rules:   rules,
		updates: updates,
	}
}

func parseLine(line, sep string) []int {
	nums := strings.Split(line, sep)
	res := make([]int, len(nums))
	for i, num := range nums {
		res[i] = lib.Atoi(num)
	}
	return res
}
