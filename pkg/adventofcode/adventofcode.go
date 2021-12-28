package adventofcode

import "context"

type Solution interface {
	String() string
}

type Puzzle interface {
	Solve(ctx context.Context) (Solution, error)
}
