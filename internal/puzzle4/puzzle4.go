package puzzle4

import (
	"errors"
	"fmt"

	"github.com/kristofferostlund/adventofcode-2021/pkg/adventofcode"
	"github.com/kristofferostlund/adventofcode-2021/pkg/fileutil"
)

func New() *adventofcode.Puzzle {
	return adventofcode.NewPuzzle(
		"puzzle 4",
		"https://adventofcode.com/2021/day/4",
		[2]int{4662, -1},
		solve,
	)
}

func solve() (answers [2]int, err error) {
	rc, err := fileutil.FileFrom("./assets/input.txt")
	if err != nil {
		return [2]int{}, fmt.Errorf("getting input: %w", err)
	}
	defer rc.Close()
	nums, boards, err := ParseInput(rc)
	if err != nil {
		return [2]int{}, fmt.Errorf("reading input: %w", err)
	}

	solution1, err := Solve1(nums, boards)
	if err != nil {
		return [2]int{}, fmt.Errorf("sovling part 1: %w", err)
	}
	solution2 := 0 // Solve2(nums, boards)

	return [2]int{solution1, solution2}, nil
}

func Solve1(nums []int, boards [][][]int) (int, error) {
	bingoBoards := FromBoards(boards)
	num, b, found := findFirstWinningBoard(nums, bingoBoards)
	if !found {
		return 0, errors.New("no winning board found")
	}

	for _, b := range bingoBoards {
		fmt.Println("")
		fmt.Println(b.TerminalDebug())
		fmt.Println("")
	}

	uncheckedSum := 0
	for _, row := range b.rows {
		for _, v := range row {
			if !b.IsChecked(v) {
				uncheckedSum += v
			}
		}
	}
	return uncheckedSum * num, nil
}
