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
		[2]int{4662, 12080},
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
	solution2, err := Solve2(nums, boards)
	if err != nil {
		return [2]int{}, fmt.Errorf("sovling part 2: %w", err)
	}

	return [2]int{solution1, solution2}, nil
}

func Solve1(nums []int, boards [][][]int) (int, error) {
	bingoBoards := FromBoards(boards)
	num, board, found := findFirstWinningBoard(nums, bingoBoards)
	if !found {
		return 0, errors.New("no winning board found")
	}

	return board.Score(num), nil
}

func Solve2(nums []int, boards [][][]int) (int, error) {
	bingoBoards := FromBoards(boards)
	num, board, found := findLastWinningBoard(nums, bingoBoards)
	if !found {
		return 0, errors.New("no winning board found")
	}

	return board.Score(num), nil
}

func findFirstWinningBoard(nums []int, boards []*BingoBoard) (int, *BingoBoard, bool) {
	for _, num := range nums {
		var found *BingoBoard
		for _, board := range boards {
			board.Check(num)
			if board.HasBingo() {
				if found == nil {
					found = board
				}
			}
		}

		if found != nil {
			return num, found, true
		}
	}
	return 0, nil, false
}

func findLastWinningBoard(nums []int, boards []*BingoBoard) (int, *BingoBoard, bool) {
	var prevBoard *BingoBoard
	prevNum := -1

	remainingBoards := make([]*BingoBoard, len(boards))
	copy(remainingBoards, boards)

	for _, num := range nums {
		indexes := make([]int, 0)

		for i, board := range remainingBoards {
			board.Check(num)
			if board.HasBingo() {
				indexes = append(indexes, i)
				prevBoard = board
				prevNum = num
			}
		}

		for i, idx := range indexes {
			copy(remainingBoards[idx-i:], remainingBoards[idx-i+1:])
			remainingBoards[len(remainingBoards)-1] = nil
			remainingBoards = remainingBoards[:len(remainingBoards)-1]
		}
	}

	return prevNum, prevBoard, prevBoard != nil
}
