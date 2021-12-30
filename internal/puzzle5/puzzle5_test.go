package puzzle5_test

import (
	"strings"
	"testing"

	"github.com/kristofferostlund/adventofcode-2021/internal/puzzle5"
	"github.com/kristofferostlund/adventofcode-2021/pkg/testhelpers"
)

func TestParseInput(t *testing.T) {
	reader := strings.NewReader(`
0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
`)

	expected := []puzzle5.Vector{
		{puzzle5.Point{0, 9}, puzzle5.Point{5, 9}},
		{puzzle5.Point{8, 0}, puzzle5.Point{0, 8}},
		{puzzle5.Point{9, 4}, puzzle5.Point{3, 4}},
	}
	actual, err := puzzle5.ParseInput(reader)
	if err != nil {
		t.Fatalf("parsing input: %s", err)
	}

	if !testhelpers.SliceEquals(actual, expected) {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}
}

func TestSolve1_exampleInput(t *testing.T) {
	input, err := puzzle5.ParseInput(strings.NewReader(`
0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2
`))
	if err != nil {
		t.Fatalf("parsing input: %s", err)
	}

	expected := 5
	actual := puzzle5.Solve1(input)
	if actual != expected {
		t.Errorf("expeted %d, got %d", expected, actual)
	}
}
