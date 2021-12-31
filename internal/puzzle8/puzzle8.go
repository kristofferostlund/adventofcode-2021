package puzzle8

import (
	"fmt"
	"io"
	"strings"

	"github.com/kristofferostlund/adventofcode-2021/pkg/adventofcode"
	"github.com/kristofferostlund/adventofcode-2021/pkg/fileutil"
)

func New() *adventofcode.Puzzle {
	return adventofcode.NewPuzzle(
		"puzzle 8",
		"https://adventofcode.com/2021/day/8",
		[2]int{470, -1},
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

type SignalPair [2][]string

func (s SignalPair) Pair() ([]string, []string) {
	return s[0], s[1]
}

func ParseInput(reader io.Reader) ([]SignalPair, error) {
	return fileutil.MapNonEmptyLines(reader, func(line string) (SignalPair, error) {
		signal, value, found := strings.Cut(line, " | ")
		if !found {
			return SignalPair{}, fmt.Errorf("could not cut line \"%s\" by \" | \"", line)
		}
		return SignalPair{strings.Split(signal, " "), strings.Split(value, " ")}, nil
	})
}

var segmentSizeToNumbers = map[int][]int{
	2: {1},
	3: {7},
	4: {4},
	5: {2, 3, 5},
	6: {0, 6, 9},
	7: {8},
}

func Solve1(input []SignalPair) int {
	count := 0
	for _, sp := range input {
		for _, segment := range sp[1] {
			if len(segmentSizeToNumbers[len(segment)]) == 1 {
				count++
			}
		}
	}

	return count
}

func Solve2(input []SignalPair) int {
	return 0
}
