package puzzle2

import (
	"github.com/kristofferostlund/adventofcode-2021/pkg/adventofcode"
)

func New() *adventofcode.Puzzle {
	return adventofcode.NewPuzzle(
		"puzzle 2",
		"https://adventofcode.com/2021/day/2",
		[2]int{-1, -1},
		solve,
	)
}

func solve() (answers [2]int, err error) {
	return [2]int{}, nil
}
