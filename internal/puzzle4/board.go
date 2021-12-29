package puzzle4

import (
	"fmt"
	"strings"

	"github.com/kristofferostlund/adventofcode-2021/pkg/stringutil"
)

type BingoBoard struct {
	rows    [][]int
	lookup  map[int][2]int
	checked map[int]map[int]struct{}
}

func NewPlayableBoard(rows [][]int) *BingoBoard {
	return &BingoBoard{
		rows:    rows,
		lookup:  createRowLookup(rows),
		checked: make(map[int]map[int]struct{}),
	}
}

func FromBoards(boards [][][]int) []*BingoBoard {
	out := make([]*BingoBoard, 0)
	for _, b := range boards {
		out = append(out, NewPlayableBoard(b))
	}
	return out
}

func createRowLookup(rows [][]int) map[int][2]int {
	lookup := make(map[int][2]int)
	for i, row := range rows {
		for j, v := range row {
			lookup[v] = [2]int{i, j}
		}
	}
	return lookup
}

func (b *BingoBoard) Check(num int) {
	loc, exists := b.lookup[num]
	if !exists {
		return
	}

	if b.checked[loc[0]] == nil {
		b.checked[loc[0]] = make(map[int]struct{})
	}
	b.checked[loc[0]][loc[1]] = struct{}{}
}

func (b *BingoBoard) IsChecked(num int) bool {
	loc, exists := b.lookup[num]
	if !exists {
		return false
	}

	_, isChecked := b.checked[loc[0]][loc[1]]
	return isChecked
}

func (b *BingoBoard) HasBingo() bool {
	for _, row := range b.checked {
		if len(row) == 5 {
			return true
		}
	}

	for i := 0; i < len(b.rows[0]); i++ {
		lineSum := 0
		for _, row := range b.rows {
			if b.IsChecked(row[i]) {
				lineSum++
			} else {
				break
			}
		}
		if lineSum == 5 {
			return true
		}
	}
	return false
}

func (b *BingoBoard) Score(num int) int {
	sum := 0
	for _, row := range b.rows {
		for _, v := range row {
			if !b.IsChecked(v) {
				sum += v
			}
		}
	}
	return sum * num
}

func (b *BingoBoard) TerminalDebug() string {
	sb := strings.Builder{}
	if b.HasBingo() {
		sb.WriteString("*** BINGO ****\n")
	}

	for i, row := range b.rows {
		if i != 0 {
			sb.WriteString("\n")
		}

		for j, v := range row {
			if j != 0 {
				sb.WriteString(" ")
			}
			s := fmt.Sprintf("%2d", v)
			if _, isChecked := b.checked[i][j]; isChecked {
				s = stringutil.Colored(s, stringutil.ColourGreen)
			}
			sb.WriteString(s)
		}
	}

	return sb.String()
}
