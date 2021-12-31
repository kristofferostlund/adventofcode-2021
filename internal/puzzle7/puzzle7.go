package puzzle7

import (
	"fmt"
	"io"
	"sort"

	"github.com/kristofferostlund/adventofcode-2021/pkg/adventofcode"
	"github.com/kristofferostlund/adventofcode-2021/pkg/fileutil"
	"github.com/kristofferostlund/adventofcode-2021/pkg/numutil"
)

func New() *adventofcode.Puzzle {
	return adventofcode.NewPuzzle(
		"puzzle 7",
		"https://adventofcode.com/2021/day/7",
		[2]int{335330, 92439766},
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

	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})

	solution1 := Solve1(input)
	solution2 := Solve2(input)

	return [2]int{solution1, solution2}, nil
}

func ParseInput(reader io.Reader) ([]int, error) {
	return fileutil.ParseSingleLineOfInts(reader)
}

func Solve1(input []int) int {
	return bestPositionOf(input, func(v int) int { return v })
}

func Solve2(input []int) int {
	return bestPositionOf(input, func(v int) int {
		cost := 0
		for i := 0; i < v; i++ {
			cost += i + 1
		}
		return cost
	})
}

func bestPositionOf(input []int, stepCostFunc func(v int) int) int {
	poses := make(map[int]int)
	for _, pos := range input {
		poses[pos] = poses[pos] + 1
	}

	bestPos := -1
	minPos, maxPos := input[0], input[len(input)-1]
	for pos := minPos; pos <= maxPos; pos++ {
		cost := 0
		for p, multiplier := range poses {
			cost += stepCostFunc(numutil.AbsInt(pos-p)) * multiplier
		}

		if cost < bestPos || bestPos == -1 {
			bestPos = cost
		}
	}

	return bestPos
}
