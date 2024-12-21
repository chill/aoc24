package lib

import "iter"

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
