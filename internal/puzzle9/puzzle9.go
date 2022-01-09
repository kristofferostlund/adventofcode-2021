package puzzle9

import (
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

	"github.com/kristofferostlund/adventofcode-2021/pkg/adventofcode"
	"github.com/kristofferostlund/adventofcode-2021/pkg/fileutil"
	"github.com/kristofferostlund/adventofcode-2021/pkg/numutil"
	"github.com/kristofferostlund/adventofcode-2021/pkg/stringutil"
)

func New() *adventofcode.Puzzle {
	return adventofcode.NewPuzzle(
		"puzzle 9",
		"https://adventofcode.com/2021/day/9",
		[2]int{532, 1110780},
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

func ParseInput(reader io.Reader) (Grid, error) {
	rows, err := fileutil.MapNonEmptyLines(reader, func(line string) ([]int, error) {
		return numutil.Atois(strings.Split(line, ""))
	})
	if err != nil {
		return nil, err
	}
	return Grid(rows), nil
}

type Grid [][]int

func (g Grid) At(p Point) (int, bool) {
	if p.X < 0 || p.X >= len(g[0]) {
		return 0, false
	}
	if p.Y < 0 || p.Y >= len(g) {
		return 0, false
	}
	return g[p.Y][p.X], true
}

func (g Grid) Surrounding(p Point) map[Point]int {
	out := make(map[Point]int)
	for y := p.Y - 1; y <= p.Y+1; y++ {
		for x := p.X - 1; x <= p.X+1; x++ {
			curr := Point{x, y}
			if p.Equal(curr) {
				continue
			}

			if v, ok := g.At(curr); ok {
				out[curr] = v
			}
		}
	}
	return out
}

type Point struct{ X, Y int }

func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func (p Point) Equal(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}

func Solve1(grid Grid) int {
	lowPoints := make(map[Point]int)

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			point := Point{x, y}
			value, _ := grid.At(point)

			isLowPoint := true
			for _, v := range grid.Surrounding(point) {
				if v <= value {
					isLowPoint = false
					break
				}
			}
			if isLowPoint {
				lowPoints[point] = value
			}
		}
	}

	totalRisk := 0
	for _, v := range lowPoints {
		totalRisk += v + 1
	}
	return totalRisk
}

func Solve2(grid Grid) int {
	basins := findBasins(grid)
	sort.Slice(basins, func(i, j int) bool {
		return len(basins[i]) > len(basins[j])
	})

	product := 0
	for _, b := range basins[:3] {
		if product == 0 {
			product = len(b)
			continue
		}
		product *= len(b)
	}

	return product
}

func findBasins(grid Grid) [][]Point {
	basins := make([][]Point, 0)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			point := Point{x, y}
			value, _ := grid.At(point)
			if value == 9 {
				continue
			}

			isLowPoint := true
			for _, v := range grid.Surrounding(point) {
				if v <= value {
					isLowPoint = false
					break
				}
			}
			if !isLowPoint {
				continue
			}

			basins = append(basins, basinFor(grid, point))
		}
	}

	return basins
}

func basinFor(grid Grid, point Point) []Point {
	lookup := map[Point]struct{}{}
	points := []Point{}
	addPoint := func(p Point) {
		if _, exists := lookup[p]; exists {
			return
		}
		lookup[p] = struct{}{}
		points = append(points, p)
	}

	addPoint(point)
	for i := 0; i < len(points); i++ {
		pt := points[i]
		for p, v := range grid.Surrounding(pt) {
			if v < 9 && (p.X == pt.X || p.Y == pt.Y) {
				addPoint(p)
			}
		}
	}
	return points
}

func highlightBasin(grid Grid, basin []Point) string {
	inBasin := make(map[Point]struct{})
	for _, p := range basin {
		inBasin[p] = struct{}{}
	}

	sb := strings.Builder{}
	for y := 0; y < len(grid); y++ {
		if y > 0 {
			sb.WriteString("\n")
		}

		for x := 0; x < len(grid[y]); x++ {
			p := Point{x, y}
			val, _ := grid.At(p)
			if _, exists := inBasin[p]; exists {
				sb.WriteString(stringutil.Colored(strconv.Itoa(val), stringutil.ColourRed))
			} else {
				sb.WriteString(strconv.Itoa(val))
			}
		}
	}

	return sb.String()
}
