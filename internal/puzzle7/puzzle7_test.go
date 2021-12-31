package puzzle7_test

import (
	"sort"
	"strings"
	"testing"

	"github.com/kristofferostlund/adventofcode-2021/internal/puzzle6"
	"github.com/kristofferostlund/adventofcode-2021/internal/puzzle7"
	"github.com/kristofferostlund/adventofcode-2021/pkg/testhelpers"
)

func TestParseInput(t *testing.T) {
	reader := strings.NewReader(`16,1,2,0,4,2,7,1,2,14`)
	expected := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	actual, err := puzzle7.ParseInput(reader)
	if err != nil {
		t.Fatalf("parsing input: %s", err)
	}

	if !testhelpers.SliceEquals(actual, expected) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestSolve1_exampleInput(t *testing.T) {
	input, err := puzzle6.ParseInput(strings.NewReader(`16,1,2,0,4,2,7,1,2,14`))
	if err != nil {
		t.Fatalf("parsing input: %s", err)
	}

	expected := 37
	actual := puzzle7.Solve1(sortedAsc(input))
	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func TestSolve2_exampleInput(t *testing.T) {
	input, err := puzzle6.ParseInput(strings.NewReader(`16,1,2,0,4,2,7,1,2,14`))
	if err != nil {
		t.Fatalf("parsing input: %s", err)
	}

	expected := 168
	actual := puzzle7.Solve2(sortedAsc(input))
	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func sortedAsc(input []int) []int {
	out := make([]int, len(input))
	copy(out, input)

	sort.Slice(out, func(i, j int) bool {
		return out[i] < out[j]
	})

	return out
}
