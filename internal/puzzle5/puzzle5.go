package puzzle5

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/kristofferostlund/adventofcode-2021/pkg/adventofcode"
	"github.com/kristofferostlund/adventofcode-2021/pkg/fileutil"
	"github.com/kristofferostlund/adventofcode-2021/pkg/intutil"
	"github.com/kristofferostlund/adventofcode-2021/pkg/stringutil"
)

type (
	Point  struct{ X, Y int }
	Vector [2]Point
)

func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func New() *adventofcode.Puzzle {
	return adventofcode.NewPuzzle(
		"puzzle 5",
		"https://adventofcode.com/2021/day/5",
		[2]int{5197, -1},
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
	solution2, err := Solve2(input)
	if err != nil {
		return [2]int{}, fmt.Errorf("solving part 2: %w", err)
	}

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
	xy, err := intutil.Atois(strings.Split(str, ","))
	if err != nil {
		return Point{}, fmt.Errorf("parsing xy: %w", err)
	}
	if len(xy) != 2 {
		return Point{}, fmt.Errorf("unexpected length %d of xy, must be exactly 2", len(xy))
	}
	return Point{xy[0], xy[1]}, nil
}

type SparseGrid struct {
	maxX, maxY int
	points     map[Point]int
}

func NewSparseGrid() *SparseGrid {
	return &SparseGrid{points: make(map[Point]int)}
}

func (sg *SparseGrid) add(p Point) {
	sg.points[p] = sg.points[p] + 1
	if sg.maxX < p.X {
		sg.maxX = p.X
	}
	if sg.maxY < p.Y {
		sg.maxY = p.Y
	}
}

func (sg *SparseGrid) addLine(v Vector) {
	fromY, toY := intutil.FromTo(v[0].Y, v[1].Y)
	fromX, toX := intutil.FromTo(v[0].X, v[1].X)

	for y := fromY; y <= toY; y++ {
		for x := fromX; x <= toX; x++ {
			sg.add(Point{x, y})
		}
	}
}

func (sg *SparseGrid) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("Size: %dx%d", sg.maxX+1, sg.maxY+1))

	for y := 0; y <= sg.maxY; y++ {
		sb.WriteString("\n")
		for x := 0; x <= sg.maxX; x++ {
			count := sg.points[Point{x, y}]
			switch count {
			case 0:
				sb.WriteString(".")
			case 1:
				sb.WriteString("1")
			default:
				sb.WriteString(stringutil.Colored(strconv.Itoa(count), stringutil.ColourGreen))
			}
		}
	}

	return sb.String()
}

func Solve1(input []Vector) int {
	grid := NewSparseGrid()

	for _, v := range input {
		if v[0].X == v[1].X || v[0].Y == v[1].Y {
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

func Solve2(input []Vector) (int, error) {
	return 0, nil
}
