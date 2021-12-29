package puzzle3

import (
	"fmt"
	"io"
	"strconv"

	"github.com/kristofferostlund/adventofcode-2021/pkg/adventofcode"
	"github.com/kristofferostlund/adventofcode-2021/pkg/fileutil"
)

func New() *adventofcode.Puzzle {
	return adventofcode.NewPuzzle(
		"puzzle 3",
		"https://adventofcode.com/2021/day/3",
		[2]int{4001724, -1},
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
		return [2]int{}, fmt.Errorf("solving part 1: %w", err)
	}

	return [2]int{solution1, solution2}, nil
}

func ParseInput(reader io.Reader) ([][]int, error) {
	return fileutil.MapLines(reader, func(line string) ([]int, error) {
		arr := make([]int, len(line))
		for i, v := range line {
			if i >= len(arr) {
				return []int{}, fmt.Errorf("unexpeted length %d of line \"%s\"", len(line), line)
			}
			val, err := strconv.Atoi(string(v))
			if err != nil {
				return []int{}, fmt.Errorf("parsing \"%s\": %w", string(v), err)
			}
			arr[i] = val
		}
		return arr, nil
	})
}

func Solve1(input [][]int) (int, error) {
	elemSize := len(input[0])
	sums := make([]int, elemSize)
	for _, row := range input {
		for j, v := range row {
			sums[j] += v
		}
	}

	gammaRate, epsilonRate := make([]rune, elemSize), make([]rune, elemSize)
	for i, v := range sums {
		switch true {
		case len(input)-v > v:
			gammaRate[i] = '1'
			epsilonRate[i] = '0'
		case len(input)-v < v:
			gammaRate[i] = '0'
			epsilonRate[i] = '1'
		default:
			return 0, fmt.Errorf("unexpected equilibrium, %d == %d", len(input)-v, v)
		}
	}

	g, err := strconv.ParseInt(string(gammaRate), 2, 0)
	if err != nil {
		return 0, fmt.Errorf("parsing gamma rate \"%s\" as base 2: %w", string(gammaRate), err)
	}
	e, err := strconv.ParseInt(string(epsilonRate), 2, 0)
	if err != nil {
		return 0, fmt.Errorf("parsing epsilon rate \"%s\" as base 2: %w", string(epsilonRate), err)
	}

	return int(g) * int(e), nil
}

func Solve2(input [][]int) (int, error) {
	return 0, nil
}
