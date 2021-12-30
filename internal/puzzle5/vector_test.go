package puzzle5_test

import (
	"fmt"
	"testing"

	"github.com/kristofferostlund/adventofcode-2021/internal/puzzle5"
	"github.com/kristofferostlund/adventofcode-2021/pkg/numutil"
)

func TestVector_Slope(t *testing.T) {
	cases := []struct {
		v        puzzle5.Vector
		expected float64
	}{
		{v: puzzle5.Vector{puzzle5.Point{1, 0}, puzzle5.Point{1, 2}}, expected: 0},
		{v: puzzle5.Vector{puzzle5.Point{0, 1}, puzzle5.Point{2, 1}}, expected: 0},
		{v: puzzle5.Vector{puzzle5.Point{0, 0}, puzzle5.Point{1, 1}}, expected: 1},
		{v: puzzle5.Vector{puzzle5.Point{2, 2}, puzzle5.Point{1, 1}}, expected: 1},
		{v: puzzle5.Vector{puzzle5.Point{0, 0}, puzzle5.Point{6, 2}}, expected: 1.0 / 3},
		{v: puzzle5.Vector{puzzle5.Point{3, 4}, puzzle5.Point{3, 9}}, expected: 0},
		{v: puzzle5.Vector{puzzle5.Point{9, 7}, puzzle5.Point{7, 9}}, expected: -1},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("%s -> %s has slope %f", test.v[0], test.v[1], test.expected), func(t *testing.T) {
			if !numutil.Float64sMatch(test.v.Slope(), test.expected) {
				t.Errorf("expected %f, got %f", test.expected, test.v.Slope())
			}
		})
	}
}
