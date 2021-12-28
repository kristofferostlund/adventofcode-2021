package puzzle1

import (
	"context"
	"fmt"
	"strconv"

	"github.com/kristofferostlund/adventofcode-2021/pkg/adventofcode"
	"github.com/kristofferostlund/adventofcode-2021/pkg/fileutil"
)

// https://adventofcode.com/2021/day/1
var expected = [2]int{1581, 1618}

type solution struct {
	answers [2]int
}

func (s *solution) String() string {
	return fmt.Sprintf(
		`
solution 1: % 5d, correct: %5t (expected %d, got %d)
solution 2: % 5d, correct: %5t (expected %d, got %d)`,
		s.answers[0], s.answers[0] == expected[0], expected[0], s.answers[0],
		s.answers[1], s.answers[1] == expected[1], expected[1], s.answers[1],
	)[1:]
}

type Puzzle struct{}

func New() *Puzzle {
	return &Puzzle{}
}

func (p *Puzzle) Solve(ctx context.Context) (adventofcode.Solution, error) {
	rc, err := fileutil.FileFrom("./assets/input.txt")
	if err != nil {
		return nil, fmt.Errorf("getting input: %w", err)
	}
	defer rc.Close()
	input, err := fileutil.MapLines(rc, func(line string) (int, error) { return strconv.Atoi(line) })
	if err != nil {
		return nil, fmt.Errorf("reading input: %w", err)
	}

	solution1 := p.solve1(input)
	solution2 := p.solve2(input)

	return &solution{[2]int{solution1, solution2}}, nil
}

func (p *Puzzle) solve1(input []int) int {
	counter := 0
	for i := 0; i < len(input)-1; i++ {
		if input[i] < input[i+1] {
			counter++
		}
	}
	return counter
}

func (p *Puzzle) solve2(input []int) int {
	counter := 0
	windows := [][]int{{}}
	for i := range input {
		for j := i; j < i+3 && j < len(input); j++ {
			if len(windows) <= j {
				windows = append(windows, []int{})
			}
			windows[i] = append(windows[i], input[j])
		}

		if len(windows[i]) != 3 || i == 0 {
			continue
		}
		if sumInts(windows[i-1]) < sumInts(windows[i]) {
			counter++
		}
	}

	return counter
}

func sumInts(ints []int) int {
	sum := 0
	for _, v := range ints {
		sum += v
	}
	return sum
}
