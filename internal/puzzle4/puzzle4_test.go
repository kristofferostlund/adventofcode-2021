package puzzle4_test

import (
	"strings"
	"testing"

	"github.com/kristofferostlund/adventofcode-2021/internal/puzzle4"
	"github.com/kristofferostlund/adventofcode-2021/pkg/testhelpers"
)

func TestPraseInput(t *testing.T) {
	reader := strings.NewReader(`
1,2,3,4,5

 1  2  3  4  5
 6  7  8  9 10
11 12 13 14 15
16 17 18 19 20
21 22 23 24 25

26 27 28 29 30
31 32 33 34 35
36 37 38 39 40
41 42 43 44 45
46 47 48 49 50
`)

	expectedNumbers := []int{1, 2, 3, 4, 5}
	expectedBoards := [][][]int{
		{
			{1, 2, 3, 4, 5},
			{6, 7, 8, 9, 10},
			{11, 12, 13, 14, 15},
			{16, 17, 18, 19, 20},
			{21, 22, 23, 24, 25},
		},
		{
			{26, 27, 28, 29, 30},
			{31, 32, 33, 34, 35},
			{36, 37, 38, 39, 40},
			{41, 42, 43, 44, 45},
			{46, 47, 48, 49, 50},
		},
	}

	actualNumbers, actualBoards, err := puzzle4.ParseInput(reader)
	if err != nil {
		t.Fatalf("parsing input: %s", err)
	}

	if !testhelpers.SliceEquals(actualNumbers, expectedNumbers) {
		t.Errorf("expected numbers %v, got %v", expectedBoards, actualBoards)
	}

	if len(actualBoards) != len(expectedBoards) {
		t.Fatalf("expected board count to be %d, got %d", len(expectedBoards), len(actualBoards))
	}
	for i, actual := range actualBoards {
		expected := expectedBoards[i]
		if len(actual) != len(expected) {
			t.Errorf("expected board row count to be %d, got %d", len(expected), len(actual))
		}
		for j, row := range actual {
			if !testhelpers.SliceEquals(row, expected[j]) {
				t.Errorf("expected %v, got %v", expected[j], row)
			}
		}
	}
}

func TestSolve1_exampleInput(t *testing.T) {
	nums, boards, err := puzzle4.ParseInput(strings.NewReader(`
7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
`))
	if err != nil {
		t.Fatalf("parsing input: %s", err)
	}

	expected := 4512
	actual, err := puzzle4.Solve1(nums, boards)
	if err != nil {
		t.Fatalf("solving puzzle: %s", err)
	}
	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func TestSolve1_verticalBingo(t *testing.T) {
	nums, boards, err := puzzle4.ParseInput(strings.NewReader(`
22,8,21,6,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19
`))
	if err != nil {
		t.Fatalf("parsing input: %s", err)
	}

	expected := 242
	actual, err := puzzle4.Solve1(nums, boards)
	if err != nil {
		t.Fatalf("solving puzzle: %s", err)
	}
	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}
