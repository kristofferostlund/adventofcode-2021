package adventofcode

import "fmt"

type Solution interface {
	String() string
}

type solution struct {
	expected, answers [2]int
}

func (s *solution) String() string {
	return fmt.Sprintf(
		`
solution 1: % 5d, correct: %5t (expected %d, got %d)
solution 2: % 5d, correct: %5t (expected %d, got %d)`,
		s.answers[0], s.answers[0] == s.expected[0], s.expected[0], s.answers[0],
		s.answers[1], s.answers[1] == s.expected[1], s.expected[1], s.answers[1],
	)[1:]
}

type SolveFunc func() (answers [2]int, err error)

type Puzzle struct {
	// TODO: Use name and url somehow in the solution?
	name      string
	url       string
	expected  [2]int
	solveFunc SolveFunc
}

func NewPuzzle(name, url string, expected [2]int, solvFunc SolveFunc) *Puzzle {
	return &Puzzle{expected: expected, solveFunc: solvFunc}
}

func (p *Puzzle) Solve() (Solution, error) {
	answers, err := p.solveFunc()
	if err != nil {
		return nil, fmt.Errorf("solving puzzle %s: %w", p.name, err)
	}
	return &solution{p.expected, answers}, nil
}
