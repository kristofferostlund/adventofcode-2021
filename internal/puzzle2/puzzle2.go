package puzzle2

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/kristofferostlund/adventofcode-2021/pkg/adventofcode"
	"github.com/kristofferostlund/adventofcode-2021/pkg/fileutil"
)

type Direction string

const (
	DirectionUp      Direction = "up"
	DirectionDown    Direction = "down"
	DirectionForward Direction = "forward"
)

type Step struct {
	Direction
	Value int
}

func New() *adventofcode.Puzzle {
	return adventofcode.NewPuzzle(
		"puzzle 2",
		"https://adventofcode.com/2021/day/2",
		[2]int{1882980, 1971232560},
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

func ParseInput(reader io.Reader) ([]Step, error) {
	return fileutil.MapLines(reader, func(line string) (Step, error) {
		dir, val, found := strings.Cut(line, " ")
		if !found {
			return Step{}, fmt.Errorf("could not cut line \"%s\" by \" \"", line)
		}
		value, err := strconv.Atoi(val)
		if err != nil {
			return Step{}, fmt.Errorf("parsing value for instruction: %w", err)
		}
		return Step{Direction: Direction(dir), Value: value}, nil
	})
}

func Solve1(input []Step) int {
	x, y := 0, 0
	for _, s := range input {
		switch s.Direction {
		case DirectionForward:
			x += s.Value
		case DirectionDown:
			y += s.Value
		case DirectionUp:
			y -= s.Value
		default:
			panic(fmt.Errorf("unsupported direction %s", s.Direction))
		}
	}
	return x * y
}

func Solve2(input []Step) int {
	x, y, aim := 0, 0, 0
	for _, s := range input {
		switch s.Direction {
		case DirectionForward:
			x += s.Value
			y += aim * s.Value
		case DirectionDown:
			aim += s.Value
		case DirectionUp:
			aim -= s.Value
		default:
			panic(fmt.Errorf("unsupported direction %s", s.Direction))
		}
	}
	return x * y
}
