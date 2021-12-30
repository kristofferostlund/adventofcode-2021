package puzzle5

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kristofferostlund/adventofcode-2021/pkg/numutil"
	"github.com/kristofferostlund/adventofcode-2021/pkg/stringutil"
)

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
	slope := v.Slope()
	if !numutil.Float64In([]float64{-1, 0, 1}, slope) {
		panic(fmt.Sprintf("illegal slope %f", slope))
	}

	fromY, toY := v[0].Y, v[1].Y
	fromX, toX := v[0].X, v[1].X

	incX, compX := loopDirection(fromX, toX)
	incY, compY := loopDirection(fromY, toY)

	for x, y := fromX, fromY; compX(x) && compY(y); x, y = x+incX, y+incY {
		sg.add(Point{x, y})
	}
}

func loopDirection(from, to int) (int, func(y int) bool) {
	switch true {
	case from < to:
		return 1, func(v int) bool { return v <= to }
	case from > to:
		return -1, func(v int) bool { return v >= to }
	default:
		return 0, func(v int) bool { return v == to }
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
