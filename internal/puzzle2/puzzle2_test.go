package puzzle2_test

import (
	"strings"
	"testing"

	"github.com/kristofferostlund/adventofcode-2021/internal/puzzle2"
)

func TestParseInput(t *testing.T) {
	reader := strings.NewReader(`
forward 1
up 2
down 3`)
	expected := []puzzle2.Step{
		{Direction: puzzle2.DirectionForward, Value: 1},
		{Direction: puzzle2.DirectionUp, Value: 2},
		{Direction: puzzle2.DirectionDown, Value: 3},
	}

	actual, err := puzzle2.ParseInput(reader)
	if err != nil {
		t.Fatalf("parsing input: %s", err)
	}

	if len(actual) != len(expected) {
		t.Fatalf("expected output length to be %d, got %d", len(expected), len(actual))
	}

	for i, s := range actual {
		if s.Direction != expected[i].Direction || s.Value != expected[i].Value {
			t.Errorf("expected step at %d to be %+v, got %+v", i, expected[i], s)
		}
	}
}

func TestSolve1_exampleInput(t *testing.T) {
	input, err := puzzle2.ParseInput(strings.NewReader(`
forward 5
down 5
forward 8
up 3
down 8
forward 2`))
	if err != nil {
		t.Fatalf("parsing input: %s", err)
	}

	expected := 150
	actual := puzzle2.Solve1(input)
	if expected != actual {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func TestSolve2_exampleInput(t *testing.T) {
	input, err := puzzle2.ParseInput(strings.NewReader(`
forward 5
down 5
forward 8
up 3
down 8
forward 2`))
	if err != nil {
		t.Fatalf("parsing input: %s", err)
	}

	expected := 900
	actual := puzzle2.Solve2(input)
	if expected != actual {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}
