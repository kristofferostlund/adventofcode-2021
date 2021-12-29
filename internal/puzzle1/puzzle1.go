package puzzle1

import (
	"fmt"
	"strconv"

	"github.com/kristofferostlund/adventofcode-2021/pkg/adventofcode"
	"github.com/kristofferostlund/adventofcode-2021/pkg/fileutil"
	"github.com/kristofferostlund/adventofcode-2021/pkg/intutil"
)

func New() *adventofcode.Puzzle {
	return adventofcode.NewPuzzle(
		"puzzle 1",
		"https://adventofcode.com/2021/day/1",
		[2]int{1581, 1618},
		solve,
	)
}

func solve() ([2]int, error) {
	rc, err := fileutil.FileFrom("./assets/input.txt")
	if err != nil {
		return [2]int{}, fmt.Errorf("getting input: %w", err)
	}
	defer rc.Close()
	input, err := fileutil.MapNonEmptyLines(rc, strconv.Atoi)
	if err != nil {
		return [2]int{}, fmt.Errorf("reading input: %w", err)
	}

	solution1 := solve1(input)
	solution2 := solve2(input)

	return [2]int{solution1, solution2}, nil
}

func solve1(input []int) int {
	counter := 0
	for i := 0; i < len(input)-1; i++ {
		if input[i] < input[i+1] {
			counter++
		}
	}
	return counter
}

func solve2(input []int) int {
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
		if intutil.SumInts(windows[i-1]) < intutil.SumInts(windows[i]) {
			counter++
		}
	}

	return counter
}
