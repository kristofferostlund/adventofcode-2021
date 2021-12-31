package puzzle6_test

import (
	"strings"
	"testing"

	"github.com/kristofferostlund/adventofcode-2021/internal/puzzle6"
	"github.com/kristofferostlund/adventofcode-2021/pkg/testhelpers"
)

func TestParseInput8(t *testing.T) {
	reader := strings.NewReader(`3,4,3,1,2`)
	expected := []int{3, 4, 3, 1, 2}

	actual, err := puzzle6.ParseInput(reader)
	if err != nil {
		t.Fatalf("parsing input: %s", err)
	}

	if !testhelpers.SliceEquals(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestSolve1_exampleInput(t *testing.T) {
	input, err := puzzle6.ParseInput(strings.NewReader(`3,4,3,1,2`))
	if err != nil {
		t.Fatalf("parsing input: %s", err)
	}

	expected := 5934
	actual := puzzle6.Solve1(input)
	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func TestSolve2_exampleInput(t *testing.T) {
	input, err := puzzle6.ParseInput(strings.NewReader(`3,4,3,1,2`))
	if err != nil {
		t.Fatalf("parsing input: %s", err)
	}

	expected := 26984457539
	actual := puzzle6.Solve2(input)
	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}
