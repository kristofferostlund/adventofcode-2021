package sliceutil_test

import (
	"testing"

	"github.com/kristofferostlund/adventofcode-2021/pkg/sliceutil"
	"github.com/kristofferostlund/adventofcode-2021/pkg/testhelpers"
)

func TestRemoveAt(t *testing.T) {
	coll := []int{0, 1, 2, 3, 4, 5}
	updated := sliceutil.RemoveAt(coll, 3)

	uExpected := []int{0, 1, 2, 4, 5}
	if testhelpers.SliceEquals(coll, updated) {
		t.Errorf("input expected not to be mutated (%v)", coll)
	}
	if !testhelpers.SliceEquals(updated, uExpected) {
		t.Errorf("expected %v, got %v", uExpected, coll)
	}
}
