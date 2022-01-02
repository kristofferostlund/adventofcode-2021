package sets

import (
	"fmt"
	"strings"
)

type Set[T comparable] map[T]struct{}

func FromRunes(s []rune) Set[rune] {
	set := make(Set[rune])
	for _, r := range s {
		set[r] = struct{}{}
	}
	return set
}

func (s Set[T]) Add(v ...T) {
	for _, k := range v {
		s[k] = struct{}{}
	}
}

func (s Set[T]) Delete(v ...T) {
	for _, k := range v {
		delete(s, k)
	}
}

func (s Set[T]) String() string {
	sb := &strings.Builder{}
	sep := ", "
	for k := range s {
		fmt.Fprintf(sb, "%s%v", sep, k)
	}

	return sb.String()[len(sep):]
}

func (s Set[T]) Values() []T {
	values := make([]T, 0, len(s))
	for k := range s {
		values = append(values, k)
	}
	return values
}

func (s Set[T]) Equals(other Set[T]) bool {
	if len(s) != len(other) {
		return false
	}
	for k := range s {
		if _, found := other[k]; !found {
			return false
		}
	}
	return true
}

func (s Set[T]) Intersection(other Set[T]) Set[T] {
	intersection := make(Set[T])
	for k := range s {
		if _, exists := other[k]; exists {
			intersection[k] = struct{}{}
		}
	}
	return intersection
}

func (s Set[T]) Difference(other Set[T]) Set[T] {
	diff := make(Set[T])
	for k := range s {
		if _, exists := other[k]; !exists {
			diff[k] = struct{}{}
		}
	}
	return diff
}

func (s Set[T]) SymmetricsDifference(other Set[T]) Set[T] {
	return s.Difference(other).Union(other.Difference(s))
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	union := make(Set[T])
	for k := range s {
		union[k] = struct{}{}
	}
	for k := range other {
		union[k] = struct{}{}
	}
	return union
}

func (s Set[T]) Contains(other Set[T]) bool {
	return len(other.Difference(s)) == 0
}

func (s Set[T]) Clone() Set[T] {
	return s.Union(make(Set[T]))
}
