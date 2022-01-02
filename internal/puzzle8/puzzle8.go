package puzzle8

import (
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

	"github.com/kristofferostlund/adventofcode-2021/pkg/adventofcode"
	"github.com/kristofferostlund/adventofcode-2021/pkg/fileutil"
	"github.com/kristofferostlund/adventofcode-2021/pkg/sets"
)

func New() *adventofcode.Puzzle {
	return adventofcode.NewPuzzle(
		"puzzle 8",
		"https://adventofcode.com/2021/day/8",
		[2]int{470, 989396},
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
	solution2, err := Solve2(input)
	if err != nil {
		return [2]int{}, fmt.Errorf("solving part 2: %w", err)
	}

	return [2]int{solution1, solution2}, nil
}

type SignalPair [2][]string

func (s SignalPair) Pair() ([]string, []string) {
	return s[0], s[1]
}

func ParseInput(reader io.Reader) ([]SignalPair, error) {
	return fileutil.MapNonEmptyLines(reader, func(line string) (SignalPair, error) {
		signal, value, found := strings.Cut(line, " | ")
		if !found {
			return SignalPair{}, fmt.Errorf("could not cut line \"%s\" by \" | \"", line)
		}
		return SignalPair{strings.Split(signal, " "), strings.Split(value, " ")}, nil
	})
}

var segmentSizeToNumbers = map[int][]int{
	2: {1},
	3: {7},
	4: {4},
	5: {2, 3, 5},
	6: {0, 6, 9},
	7: {8},
}

func Solve1(input []SignalPair) int {
	count := 0
	for _, sp := range input {
		for _, v := range sp[1] {
			if len(segmentSizeToNumbers[len(v)]) == 1 {
				count++
			}
		}
	}

	return count
}

func Solve2(input []SignalPair) (int, error) {
	sum := 0
	for _, sp := range input {
		val, err := interpretSignalPair(sp)
		if err != nil {
			return 0, fmt.Errorf("parsing signal pair: %w", err)
		}
		sum += val
	}
	return sum, nil
}

type sigMap struct {
	sigToNum map[string]int
	numToSig map[int]string
}

func (s sigMap) add(sig string, num int) {
	s.sigToNum[s.key(sig)] = num
	s.numToSig[num] = sig
}

func (s sigMap) signalFor(num int) (string, bool) {
	sig, exists := s.numToSig[num]
	return sig, exists
}

func (s sigMap) numFor(sig string) (int, bool) {
	num, exists := s.sigToNum[s.key(sig)]
	return num, exists
}

func (s sigMap) key(sig string) string {
	r := []rune(sig)
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	return string(r)
}

func interpretSignalPair(sp SignalPair) (int, error) {
	signal, values := sp.Pair()

	known := sigMap{sigToNum: make(map[string]int), numToSig: make(map[int]string)}
	unknown := make(map[int][]string)

	for _, v := range signal {
		numbers := segmentSizeToNumbers[len(v)]
		if len(numbers) == 1 {
			known.add(v, numbers[0])
		} else {
			unknown[len(v)] = append(unknown[len(v)], v)
		}
	}

	oneSig, exists := known.signalFor(1)
	if !exists {
		return 0, fmt.Errorf("no signal for %d", 1)
	}
	one := sets.FromRunes([]rune(oneSig))
	fourSig, exists := known.signalFor(4)
	if !exists {
		return 0, fmt.Errorf("no signal for %d", 4)
	}
	four := sets.FromRunes([]rune(fourSig))

	for size, signals := range unknown {
		var mappedSignals map[string]int
		switch size {
		case 5:
			mappedSignals = get235(signals, one, four)
		case 6:
			mappedSignals = get069(signals, one, four)
		default:
			return 0, fmt.Errorf("unknown segment size %d (%v)", size, signals)
		}
		if len(mappedSignals) != len(signals) {
			return 0, fmt.Errorf("unexpected output when getting %v. Got %d mapped signals (%v), expected %d", segmentSizeToNumbers[size], len(mappedSignals), mappedSignals, len(signals))
		}
		for sig, num := range mappedSignals {
			known.add(sig, num)
		}
	}

	return parseValue(values, known)
}

func get235(signals []string, one, four sets.Set[rune]) map[string]int {
	out := make(map[string]int)
	for _, sig := range signals {
		s := sets.FromRunes([]rune(sig))
		switch true {
		case s.Contains(one):
			out[sig] = 3
			continue
		case s.Union(one).Contains(four):
			out[sig] = 5
			continue
		default:
			out[sig] = 2
		}
	}
	return out
}

func get069(signals []string, one, four sets.Set[rune]) map[string]int {
	out := make(map[string]int)
	for _, sig := range signals {
		s := sets.FromRunes([]rune(sig))
		switch true {
		case len(one.Difference(s)) == 1:
			out[string(s.Values())] = 6
		case s.Contains(four):
			out[string(s.Values())] = 9
		default:
			out[string(s.Values())] = 0
		}
	}
	return out
}

func parseValue(values []string, known sigMap) (int, error) {
	sb := strings.Builder{}
	for _, v := range values {
		val, found := known.numFor(v)
		if !found {
			return 0, fmt.Errorf("unknown value segment \"%s\"", v)
		}
		sb.WriteString(strconv.Itoa(val))
	}
	return strconv.Atoi(sb.String())
}
