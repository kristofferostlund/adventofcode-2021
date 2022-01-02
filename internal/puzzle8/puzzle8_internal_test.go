package puzzle8

import "testing"

func Test_interpretSignalPair(t *testing.T) {
	sp := SignalPair{
		[]string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"},
		[]string{"cdfeb", "fcadb", "cdfeb", "cdbaf"},
	}

	expected := 5353
	actual, err := interpretSignalPair(sp)
	if err != nil {
		t.Fatalf("parsing signal pair: %s", err)
	}

	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}
