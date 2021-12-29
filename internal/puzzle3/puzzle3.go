package puzzle3

import (
	"fmt"
	"io"
	"strconv"

	"github.com/kristofferostlund/adventofcode-2021/pkg/adventofcode"
	"github.com/kristofferostlund/adventofcode-2021/pkg/fileutil"
	"github.com/kristofferostlund/adventofcode-2021/pkg/stringutil"
)

func New() *adventofcode.Puzzle {
	return adventofcode.NewPuzzle(
		"puzzle 3",
		"https://adventofcode.com/2021/day/3",
		[2]int{4001724, 587895},
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
		return [2]int{}, fmt.Errorf("solving part 2: %w", err)
	}

	return [2]int{solution1, solution2}, nil
}

func ParseInput(reader io.Reader) ([][]int, error) {
	return fileutil.MapNonEmptyLines(reader, func(line string) ([]int, error) {
		arr := make([]int, len(line))
		for i, v := range line {
			if i >= len(arr) {
				return []int{}, fmt.Errorf("unexpected length %d of line \"%s\"", len(line), line)
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
	width := len(input[0])
	sums := make([]int, width)
	for _, row := range input {
		for j, v := range row {
			sums[j] += v
		}
	}

	gammaRate, epsilonRate := make([]int, width), make([]int, width)
	for i, v := range sums {
		switch true {
		case len(input)-v > v:
			gammaRate[i] = 1
			epsilonRate[i] = 0
		case len(input)-v < v:
			gammaRate[i] = 0
			epsilonRate[i] = 1
		default:
			return 0, fmt.Errorf("unexpected equilibrium, %d == %d", len(input)-v, v)
		}
	}

	g, err := intSliceBinaryToDecimal(gammaRate)
	if err != nil {
		return 0, fmt.Errorf("parsing gamma rate %w", err)
	}
	e, err := intSliceBinaryToDecimal(epsilonRate)
	if err != nil {
		return 0, fmt.Errorf("parsing epsilon rate %w", err)
	}

	return g * e, nil
}

func Solve2(input [][]int) (int, error) {
	width := len(input[0])

	oxElems := input
	for x := 0; x < width; x++ {
		outcomes := make(map[int][][]int)
		for _, v := range oxElems {
			outcomes[v[x]] = append(outcomes[v[x]], v)
		}
		if len(outcomes[0]) > len(outcomes[1]) {
			oxElems = outcomes[0]
		} else {
			oxElems = outcomes[1]
		}
	}

	co2Elems := input
	for x := 0; x < width; x++ {
		if len(co2Elems) == 1 {
			break
		}
		outcomes := make(map[int][][]int)
		for _, v := range co2Elems {
			outcomes[v[x]] = append(outcomes[v[x]], v)
		}
		if len(outcomes[0]) <= len(outcomes[1]) {
			co2Elems = outcomes[0]
		} else {
			co2Elems = outcomes[1]
		}
	}

	if len(oxElems) != 1 {
		return 0, fmt.Errorf("unexpected number of oxygen generator readings %d, exepcted %d", len(oxElems), 1)
	}
	oxGenRating, err := intSliceBinaryToDecimal(oxElems[0])
	if err != nil {
		return 0, fmt.Errorf("parsing oxygen generator reading: %w", err)
	}

	if len(co2Elems) != 1 {
		return 0, fmt.Errorf("unexpected number of CO2 generator readings %d, exepcted %d", len(co2Elems), 1)
	}
	co2GenRating, err := intSliceBinaryToDecimal(co2Elems[0])
	if err != nil {
		return 0, fmt.Errorf("parsing CO2 generator reading: %w", err)
	}

	return oxGenRating * co2GenRating, nil
}

func intSliceBinaryToDecimal(input []int) (int, error) {
	binaryStr := stringutil.JoinAny(input, "%d", "")
	g, err := strconv.ParseInt(binaryStr, 2, 0)
	if err != nil {
		return 0, fmt.Errorf("parsing binary string \"%s\": %w", binaryStr, err)
	}
	return int(g), nil
}
