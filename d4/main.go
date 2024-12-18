package main

import (
	"fmt"

	"aoc24/lib"
)

func main() {
	lines := lib.ByLines("./input.txt")
	matrix := lib.BuildMatrix(lines, lib.Runes)
	words := xmasSearch(matrix)
	crosses := crossSearch(matrix)
	fmt.Printf("got %d occurences of XMAS, %d occurences of X-MAS\n", words, crosses)
}

func xmasSearch(matrix [][]rune) int {
	return searchByChar(matrix, 'X', checkXmas)
}

func crossSearch(matrix [][]rune) int {
	return searchByChar(matrix, 'A', checkCross)
}

type checkFunc func(y, x int, matrix [][]rune) int

func searchByChar(matrix [][]rune, target rune, check checkFunc) int {
	var count int
	for y := 0; y < len(matrix); y++ {
		row := matrix[y]
		for x := 0; x < len(row); x++ {
			if matrix[y][x] != target {
				continue
			}

			count += check(y, x, matrix)
		}
	}
	return count
}

func checkCross(y, x int, matrix [][]rune) int {
	// from an A, we only need to check 4 positions, 2 types of slash
	// backslash: 1,-1 and -1,1
	// forward slash: 1,1 and -1,-1
	// we need them all to be only M or S
	// we need the inverse of each coordinate diff to be the other letter
	// 1,-1 being M means -1,1 must be S, or vice versa
	var iter int
	var expectFwd, expectBack rune
	for diffY := -1; diffY < 2; diffY += 2 {
		for diffX := -1; diffX < 2; diffX += 2 {
			newY := y + diffY
			newX := x + diffX

			if !lib.InBounds(newY, newX, matrix) {
				return 0
			}

			var expect rune
			got := matrix[newY][newX]
			switch got {
			case 'M':
				expect = 'S'
			case 'S':
				expect = 'M'
			default:
				return 0
			}

			// order:
			// -1,-1 0
			// -1,1, 1
			// 1,-1, 2
			// 1,1,  3
			switch iter {
			case 0:
				expectFwd = expect
			case 1:
				expectBack = expect
			case 2:
				if expectBack != got {
					return 0
				}
			case 3:
				if expectFwd != got {
					return 0
				}
			default:
				panic("more iterations than expected")
			}

			iter++
		}
	}

	return 1
}

var searchXmas = []rune{'M', 'A', 'S'}

func checkXmas(y, x int, matrix [][]rune) int {
	var count int
	for diffY := -1; diffY < 2; diffY++ {
		for diffX := -1; diffX < 2; diffX++ {
			if checkLine(y, x, diffY, diffX, matrix, searchXmas) {
				count++
			}
		}
	}

	return count
}

func checkLine(y, x, diffY, diffX int, matrix [][]rune, target []rune) bool {
	for i, r := range target {
		newY := y + diffY*(i+1)
		newX := x + diffX*(i+1)
		if !lib.InBounds(newY, newX, matrix) {
			return false
		}

		if matrix[newY][newX] != r {
			return false
		}
	}

	return true
}
