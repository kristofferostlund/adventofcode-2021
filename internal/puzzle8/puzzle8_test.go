package puzzle8_test

import (
	"strings"
	"testing"

	"github.com/kristofferostlund/adventofcode-2021/internal/puzzle8"
	"github.com/kristofferostlund/adventofcode-2021/pkg/testhelpers"
)

func TestParseInput(t *testing.T) {
	reader := strings.NewReader(`
be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
`)

	expected := []puzzle8.SignalPair{
		{[]string{"be", "cfbegad", "cbdgef", "fgaecd", "cgeb", "fdcge", "agebfd", "fecdb", "fabcd", "edb"}, []string{"fdgacbe", "cefdb", "cefbgd", "gcbe"}},
		{[]string{"edbfga", "begcd", "cbg", "gc", "gcadebf", "fbgde", "acbgfd", "abcde", "gfcbed", "gfec"}, []string{"fcgedb", "cgb", "dgebacf", "gc"}},
	}
	actual, err := puzzle8.ParseInput(reader)
	if err != nil {
		t.Fatalf("parsing input: %s", err)
	}

	if len(actual) != len(expected) {
		t.Fatalf("expected length %d, got %d", len(expected), len(actual))
	}

	for i, v := range actual {
		if !testhelpers.SliceEquals(v[0], expected[i][0]) {
			t.Errorf("expected signal at %d to match %v, got %v", i, v[0], expected[i][0])
		}
		if !testhelpers.SliceEquals(v[1], expected[i][1]) {
			t.Errorf("expected value at %d to match %v, got %v", i, v[1], expected[i][1])
		}
	}
}

func TestSolve1_exampleInput(t *testing.T) {
	input, err := puzzle8.ParseInput(strings.NewReader(`
be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce
`))
	if err != nil {
		t.Fatalf("parsing input: %s", err)
	}

	expected := 26
	actual := puzzle8.Solve1(input)
	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func TestSolve2_exampleInput(t *testing.T) {
	input, err := puzzle8.ParseInput(strings.NewReader(`
be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce
`))
	if err != nil {
		t.Fatalf("parsing input: %s", err)
	}

	expected := 61229
	actual, err := puzzle8.Solve2(input)
	if err != nil {
		t.Fatalf("solving puzzle: %s", err)
	}
	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}
