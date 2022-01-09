package puzzle9_test

import (
	"strings"
	"testing"

	"github.com/kristofferostlund/adventofcode-2021/internal/puzzle9"
	"github.com/kristofferostlund/adventofcode-2021/pkg/testhelpers"
)

func TestParseInput(t *testing.T) {
	reader := strings.NewReader(`
2199943210
3987894921
9856789892
8767896789
9899965678
`)
	expected := puzzle9.Grid{
		[]int{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		[]int{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		[]int{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		[]int{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		[]int{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}
	actual, err := puzzle9.ParseInput(reader)
	if err != nil {
		t.Fatalf("parsing input: %s", err)
	}

	if len(actual) != len(expected) {
		t.Fatalf("expected %d rows, got %d", len(expected), len(actual))
	}
	for i, row := range actual {
		if !testhelpers.SliceEquals(row, expected[i]) {
			t.Errorf("expected row %d to be %v, got %v", i, expected[i], row)
		}
	}
}

func TestSolve1_exampleInput(t *testing.T) {
	input, err := puzzle9.ParseInput(strings.NewReader(`
2199943210
3987894921
9856789892
8767896789
9899965678
`))
	if err != nil {
		t.Fatalf("parsing input: %s", err)
	}
	expected := 15
	actual := puzzle9.Solve1(input)
	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func TestSolve2_exampleInput(t *testing.T) {
	input, err := puzzle9.ParseInput(strings.NewReader(`
2199943210
3987894921
9856789892
8767896789
9899965678
`))
	if err != nil {
		t.Fatalf("parsing input: %s", err)
	}
	expected := 1134
	actual := puzzle9.Solve2(input)
	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}
