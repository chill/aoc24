package lib

import (
	"iter"
	"maps"
	"slices"
)

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
