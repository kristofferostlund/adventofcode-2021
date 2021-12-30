package numutil

import (
	"fmt"
	"math"
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

func Float64sMatch(a, b float64) bool {
	return math.Abs(a-b) < 1e-6
}

func Float64In(ff []float64, f float64) bool {
	for _, fl := range ff {
		if Float64sMatch(fl, f) {
			return true
		}
	}
	return false
}
