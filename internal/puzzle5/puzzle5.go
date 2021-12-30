package puzzle5

import (
	"fmt"
	"io"
	"strings"

	"github.com/kristofferostlund/adventofcode-2021/pkg/adventofcode"
	"github.com/kristofferostlund/adventofcode-2021/pkg/fileutil"
	"github.com/kristofferostlund/adventofcode-2021/pkg/numutil"
)

type Point struct{ X, Y int }

func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

type Vector [2]Point

func (v Vector) Slope() float64 {
	y := float64(v[1].Y) - float64(v[0].Y)
	x := float64(v[1].X) - float64(v[0].X)

	if x == 0 {
		// Slope of of a horizontal line is undefined
		return 0
	}
	return y / x
}

func New() *adventofcode.Puzzle {
	return adventofcode.NewPuzzle(
		"puzzle 5",
		"https://adventofcode.com/2021/day/5",
		[2]int{5197, 18605},
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

func ParseInput(reader io.Reader) ([]Vector, error) {
	return fileutil.MapNonEmptyLines(reader, func(line string) (Vector, error) {
		p1s, p2s, found := strings.Cut(line, " -> ")
		if !found {
			return Vector{}, fmt.Errorf("malformed line \"%s\"", line)
		}
		p1, err := parsePoint(p1s)
		if err != nil {
			return Vector{}, fmt.Errorf("parsing point 1: %w", err)
		}
		p2, err := parsePoint(p2s)
		if err != nil {
			return Vector{}, fmt.Errorf("parsing point 2: %w", err)
		}
		return Vector{p1, p2}, nil
	})
}

func parsePoint(str string) (Point, error) {
	xy, err := numutil.Atois(strings.Split(str, ","))
	if err != nil {
		return Point{}, fmt.Errorf("parsing xy: %w", err)
	}
	if len(xy) != 2 {
		return Point{}, fmt.Errorf("unexpected length %d of xy, must be exactly 2", len(xy))
	}
	return Point{xy[0], xy[1]}, nil
}

func Solve1(input []Vector) int {
	grid := NewSparseGrid()

	for _, v := range input {
		if numutil.Float64sMatch(v.Slope(), 0) {
			grid.addLine(v)
		}
	}

	score := 0
	for _, count := range grid.points {
		if count >= 2 {
			score++
		}
	}

	return score
}

func Solve2(input []Vector) int {
	grid := NewSparseGrid()

	for _, v := range input {
		if numutil.Float64In([]float64{0, 1, -1}, v.Slope()) {
			grid.addLine(v)
		}
	}

	score := 0
	for _, count := range grid.points {
		if count >= 2 {
			score++
		}
	}
	return score
}
