package puzzle4

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/kristofferostlund/adventofcode-2021/pkg/fileutil"
)

func ParseInput(reader io.Reader) ([]int, [][][]int, error) {
	numbers := make([]int, 0)
	boards := make([][][]int, 0)
	board := make([][]int, 0)

	ingestNumbers := func(line string) error {
		var err error
		numbers, err = parseNumbers(strings.Split(line, ","))
		if err != nil {
			return fmt.Errorf("parsing number line: %w", err)
		}
		return nil
	}

	ingestBoards := func(line string) error {
		nums, err := parseNumbers(splitBoardLine(line))
		if err != nil {
			return fmt.Errorf("parsing board line: %w", err)
		}
		board = append(board, nums)
		if len(board) == 5 {
			boards = append(boards, board)
			board = make([][]int, 0)
		}
		return nil
	}

	if err := fileutil.ScanLines(reader, func(_ int, line string) error {
		if line == "" {
			return nil
		}
		if len(numbers) == 0 {
			return ingestNumbers(line)
		}
		return ingestBoards(line)
	}); err != nil {
		return nil, nil, fmt.Errorf("parsing lines: %w", err)
	}

	return numbers, boards, nil
}

func splitBoardLine(line string) []string {
	out := make([]string, 0)
	sb := strings.Builder{}
	for i, r := range line {
		if r == ' ' {
			continue
		}
		sb.WriteRune(r)
		if len(line) == i+1 || line[i+1] == ' ' {
			out = append(out, sb.String())
			sb.Reset()
		}
	}
	return out
}

func parseNumbers(in []string) ([]int, error) {
	out := make([]int, 0)
	for _, s := range in {
		v, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("parsing \"%s\": %w", s, err)
		}
		out = append(out, v)
	}
	return out, nil
}
