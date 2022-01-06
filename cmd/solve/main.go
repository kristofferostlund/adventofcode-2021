package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/kristofferostlund/adventofcode-2021/internal/puzzle1"
	"github.com/kristofferostlund/adventofcode-2021/internal/puzzle2"
	"github.com/kristofferostlund/adventofcode-2021/internal/puzzle3"
	"github.com/kristofferostlund/adventofcode-2021/internal/puzzle4"
	"github.com/kristofferostlund/adventofcode-2021/internal/puzzle5"
	"github.com/kristofferostlund/adventofcode-2021/internal/puzzle6"
	"github.com/kristofferostlund/adventofcode-2021/internal/puzzle7"
	"github.com/kristofferostlund/adventofcode-2021/internal/puzzle8"
	"github.com/kristofferostlund/adventofcode-2021/internal/puzzle9"
	"github.com/kristofferostlund/adventofcode-2021/pkg/adventofcode"
)

var (
	puzzle  = flag.String("puzzle", "", "what day to solve? Example: 1")
	puzzles = map[string]adventofcode.Puzzle{
		"1": *puzzle1.New(),
		"2": *puzzle2.New(),
		"3": *puzzle3.New(),
		"4": *puzzle4.New(),
		"5": *puzzle5.New(),
		"6": *puzzle6.New(),
		"7": *puzzle7.New(),
		"8": *puzzle8.New(),
		"9": *puzzle9.New(),
	}
)

func main() {
	flag.Parse()

	toSolve, ok := puzzles[*puzzle]
	if !ok {
		log.Fatalf("No such puzzle \"%s\"", *puzzle)
	}

	solution, err := toSolve.Solve()
	if err != nil {
		log.Fatalf("Solving puzzle %s: %s", *puzzle, err)
	}

	log.Printf("Attempted solving puzzle %s", *puzzle)
	fmt.Println(solution)
}
