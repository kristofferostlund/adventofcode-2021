package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/kristofferostlund/adventofcode-2021/internal/puzzle1"
	"github.com/kristofferostlund/adventofcode-2021/pkg/adventofcode"
)

var (
	puzzle  = flag.String("puzzle", "", "what day to solve? Example: 1")
	puzzles = map[string]adventofcode.Puzzle{
		"1": puzzle1.New(),
	}
)

func main() {
	flag.Parse()

	toSolve, ok := puzzles[*puzzle]
	if !ok {
		log.Fatalf("no such puzzle \"%s\"", *puzzle)
	}

	ctx := context.Background()
	solution, err := toSolve.Solve(ctx)
	if err != nil {
		log.Fatalf("solving puzzle %s: %s", *puzzle, err)
	}

	log.Printf("Attempted solving puzzle %s", *puzzle)
	fmt.Println(solution)
}
