package puzzle6

import (
	"fmt"
	"io"

	"github.com/kristofferostlund/adventofcode-2021/pkg/adventofcode"
	"github.com/kristofferostlund/adventofcode-2021/pkg/fileutil"
)

func New() *adventofcode.Puzzle {
	return adventofcode.NewPuzzle(
		"puzzle 6",
		"https://adventofcode.com/2021/day/6",
		[2]int{386536, 1732821262171},
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

func ParseInput(reader io.Reader) ([]int, error) {
	return fileutil.ParseSingleLineOfInts(reader)
}

func Solve1(input []int) int {
	fish := make([]int, len(input))
	copy(fish, input)

	for day := 0; day < 80; day++ {
		spawned := make([]int, 0)
		for i, f := range fish {
			fish[i] = f - 1
			if f-1 == -1 {
				fish[i] = 6
				spawned = append(spawned, 8)
			}
		}
		fish = append(fish, spawned...)
	}

	return len(fish)
}

func Solve2(input []int) int {
	fish := make([][2]int, 0)
	for _, f := range input {
		fish = append(fish, [2]int{f, 1})
	}

	for day := 0; day < 256; day++ {
		spawned := 0
		for i, f := range fish {
			fish[i][0] = f[0] - 1
			if f[0]-1 == -1 {
				fish[i][0] = 6
				spawned += fish[i][1]
			}
		}

		if spawned > 0 {
			fish = append(fish, [2]int{8, spawned})
		}
	}

	count := 0
	for _, f := range fish {
		count += f[1]
	}

	return count
}
