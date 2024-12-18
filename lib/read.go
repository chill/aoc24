package lib

import (
	"bufio"
	"io"
	"iter"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"
)

func ByLines(path string) iter.Seq[string] {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	return scan(f, bufio.ScanLines, stringer)
}

func StringLines(s string) iter.Seq[string] {
	return scan(strings.NewReader(s), bufio.ScanLines, stringer)
}

func Lines(path string) []string {
	return slices.Collect(ByLines(path))
}

func Words(line string) []string {
	return slices.Collect(scan(strings.NewReader(line), bufio.ScanWords, stringer))
}

func Runes(line string) []rune {
	return slices.Collect(slices.Values([]rune(line)))
}

func IntsWords(line string) []int {
	return slices.Collect(scan(strings.NewReader(line), bufio.ScanWords, inter))
}

func IntsSlice(words []string) []int {
	return ConvSlice(words, inter)
}

func ConvSlice[V any, B any](in []V, conv func(V) B) []B {
	res := make([]B, len(in))
	for i, val := range in {
		res[i] = conv(val)
	}
	return res
}

func scan[V any](r io.Reader, splitFunc bufio.SplitFunc, convert func(s string) V) iter.Seq[V] {
	scanner := bufio.NewScanner(r)
	scanner.Split(splitFunc)

	return func(yield func(V) bool) {
		for scanner.Scan() {
			if !yield(convert(scanner.Text())) {
				return
			}
		}
		if scanner.Err() != nil {
			panic(scanner.Err())
		}
	}
}

func stringer(s string) string {
	return s
}

func inter(s string) int {
	return Atoi(s)
}

func Atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}

func Abs(i int) int {
	if i < 0 {
		return i * -1
	}

	return i
}

func SafeDel[V any](vs []V, i int) []V {
	res := make([]V, len(vs)-1)
	n := copy(res, vs[:i])
	copy(res[n:], vs[i+1:])
	return res
}

func InBounds[G any](y, x int, m [][]G) bool {
	return y >= 0 && x >= 0 && y < len(m) && x < len(m[0])
}

func BuildMatrix[V any](lines iter.Seq[string], conv func(string) []V) [][]V {
	matrix := make([][]V, 0)
	for line := range lines {
		matrix = append(matrix, conv(line))
	}

	return matrix
}

func ApplyMatrix[V any](m [][]V, fn func(V) V) {
	for _, r := range m {
		for i, v := range r {
			r[i] = fn(v)
		}
	}
}

type Set[V comparable] map[V]struct{}

func NewSet[V comparable](vals ...V) Set[V] {
	s := Set[V]{}
	s.Add(vals...)
	return s
}

func (s Set[V]) Add(vals ...V) {
	for _, v := range vals {
		s[v] = struct{}{}
	}
}

func (s Set[V]) Contains(vals ...V) bool {
	contained := true
	for _, v := range vals {
		_, ok := s[v]
		contained = contained && ok
	}

	return contained
}

func (s Set[V]) Delete(vals ...V) {
	for _, v := range vals {
		delete(s, v)
	}
}

func (s Set[V]) Equals(other Set[V]) bool {
	if len(s) != len(other) {
		return false
	}

	return s.Contains(slices.Collect(maps.Keys(other))...)
}

func (s Set[V]) Values() iter.Seq[V] {
	return maps.Keys(s)
}
