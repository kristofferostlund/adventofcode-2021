package puzzle3_test

import (
	"strings"
	"testing"

	"github.com/kristofferostlund/adventofcode-2021/internal/puzzle3"
	"github.com/kristofferostlund/adventofcode-2021/pkg/testhelpers"
)

func TestParseInput(t *testing.T) {
	reader := strings.NewReader(`
00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`)

	expected := [][]int{
		{0, 0, 1, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 1, 1, 0},
		{1, 0, 1, 1, 1},
		{1, 0, 1, 0, 1},
		{0, 1, 1, 1, 1},
		{0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 1},
		{0, 0, 0, 1, 0},
		{0, 1, 0, 1, 0},
	}
	actual, err := puzzle3.ParseInput(reader)
	if err != nil {
		t.Fatalf("parsing input: %s", err)
	}

	if len(actual) != len(expected) {
		t.Fatalf("expected output length to be %d, got %d", len(expected), len(actual))
	}

	for i, v := range actual {
		if !testhelpers.SliceEquals(v, expected[i]) {
			t.Errorf("expected element at %d to match %v, got %v", i, expected[i], v)
		}
	}
}

func TestSolve1_exampleInput(t *testing.T) {
	input, err := puzzle3.ParseInput(strings.NewReader(`
00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`))
	if err != nil {
		t.Fatalf("parsing input: %s", err)
	}

	expected := 198
	actual, err := puzzle3.Solve1(input)
	if err != nil {
		t.Fatalf("solving puzzle: %s", err)
	}
	if expected != actual {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func TestSolve2_exampleInput(t *testing.T) {
	input, err := puzzle3.ParseInput(strings.NewReader(`
00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`))
	if err != nil {
		t.Fatalf("parsing input: %s", err)
	}

	expected := 230
	actual, err := puzzle3.Solve2(input)
	if err != nil {
		t.Fatalf("solving puzzle: %s", err)
	}
	if expected != actual {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}
