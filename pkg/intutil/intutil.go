package intutil

import (
	"fmt"
	"strconv"
)

func SumInts(ints []int) int {
	sum := 0
	for _, v := range ints {
		sum += v
	}
	return sum
}

func Atois(in []string) ([]int, error) {
	out := make([]int, 0)
	for _, s := range in {
		v, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("parsing \"%s\": %w", s, err)
		}
		out = append(out, v)
	}
	return out, nil
}

func FromTo(a, b int) (int, int) {
	if a <= b {
		return a, b
	}
	return b, a
}
