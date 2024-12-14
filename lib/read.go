package lib

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func ByLines(path string) <-chan string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	return scan(f, bufio.ScanLines, stringer)
}

func Lines(path string) []string {
	return acc(ByLines(path))
}

func Words(line string) []string {
	return acc(scan(strings.NewReader(line), bufio.ScanWords, stringer))
}

func Ints(line string) []int {
	return acc(scan(strings.NewReader(line), bufio.ScanWords, inter))
}

func scan[V any](r io.Reader, splitFunc bufio.SplitFunc, convert func(s string) V) <-chan V {
	scanner := bufio.NewScanner(r)
	scanner.Split(splitFunc)

	res := make(chan V, 10)
	go func() {
		for scanner.Scan() {
			res <- convert(scanner.Text())
		}
		if scanner.Err() != nil {
			panic(scanner.Err())
		}
		close(res)
	}()

	return res
}

func stringer(s string) string {
	return s
}

func inter(s string) int {
	return Atoi(s)
}

func acc[V any](ch <-chan V) []V {
	res := make([]V, 0)
	for e := range ch {
		res = append(res, e)
	}
	return res
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
