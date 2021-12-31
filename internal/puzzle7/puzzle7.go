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
		[2]int{335330, -1},
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

	solution1 := Solve1(input)
	solution2 := Solve2(input)

	return [2]int{solution1, solution2}, nil
}

func ParseInput(reader io.Reader) ([]int, error) {
	return fileutil.ParseSingleLineOfInts(reader)
}

func Solve1(input []int) int {
	crabs := make([]int, len(input))
	copy(crabs, input)
	sort.Slice(crabs, func(i, j int) bool {
		return crabs[i] < crabs[j]
	})

	poses := make(map[int]int)
	for _, pos := range crabs {
		poses[pos] = poses[pos] + 1
	}

	minPos, maxPos := crabs[0], crabs[len(crabs)-1]
	bestPos := -1
	for pos := minPos; pos <= maxPos; pos++ {
		cost := 0
		for p, multiplier := range poses {
			cost += numutil.AbsInt(pos-p) * multiplier
		}
		if cost < bestPos || bestPos == -1 {
			bestPos = cost
		}
	}

	return bestPos
}

func Solve2(input []int) int {
	return 0
}
