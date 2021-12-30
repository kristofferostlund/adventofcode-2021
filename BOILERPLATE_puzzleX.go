package puzzleX

import (
	"fmt"
	"io"

	"github.com/kristofferostlund/adventofcode-2021/pkg/adventofcode"
	"github.com/kristofferostlund/adventofcode-2021/pkg/fileutil"
)

func New() *adventofcode.Puzzle {
	return adventofcode.NewPuzzle(
		"puzzle X",
		"https://adventofcode.com/2021/day/X",
		[2]int{-1, -1},
		solve,
	)
}

func solve() (answers [2]int, err error) {
	rc, err := fileutil.FileFrom("./assets/input.txt")
	if err != nil {
		return [2]int{}, fmt.Errorf("getting input: %w", err)
	}
	defer rc.Close()
	input, err := ParseInput(rc)
	if err != nil {
		return [2]int{}, fmt.Errorf("reading input: %w", err)
	}

	solution1, err := Solve1(input)
	if err != nil {
		return [2]int{}, fmt.Errorf("solving part 1: %w", err)
	}
	solution2, err := Solve2(input)
	if err != nil {
		return [2]int{}, fmt.Errorf("solving part 2: %w", err)
	}

	return [2]int{solution1, solution2}, nil
}

func ParseInput(reader io.Reader) ([]any, error) {
	return fileutil.MapNonEmptyLines(reader, func(line string) (any, error) {
		return nil, nil
	})
}

func Solve1(input []any) (int, error) {
	return 0, nil
}

func Solve2(input []any) (int, error) {
	return 0, nil
}
