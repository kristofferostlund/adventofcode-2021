package puzzle2_test

import (
	"strings"
	"testing"

	"github.com/kristofferostlund/adventofcode-2021/internal/puzzle2"
)

func TestParseInput(t *testing.T) {
	input := strings.NewReader(`
forward 1
up 2
down 3`[1:])
	expected := []puzzle2.Step{
		{Direction: puzzle2.DirectionForward, Value: 1},
		{Direction: puzzle2.DirectionUp, Value: 2},
		{Direction: puzzle2.DirectionDown, Value: 3},
	}

	output, err := puzzle2.ParseInput(input)
	if err != nil {
		t.Fatalf("parsing input: %s", err)
	}

	if len(output) != len(expected) {
		t.Fatalf("expected output length to be %d, got %d", len(expected), len(output))
	}

	for i, s := range output {
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
forward 2`[1:]))
	if err != nil {
		t.Fatalf("parsing input: %s", err)
	}

	expected := 150
	output := puzzle2.Solve1(input)
	if expected != output {
		t.Errorf("expected %d, got %d", expected, output)
	}
}
