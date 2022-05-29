package day04

import (
	"testing"

	"github.com/imaspen/aoc-2015-go/inputreader"
)

func TestPart1(t *testing.T) {
	expected := [...]int{609043, 1048970}
	lines := inputreader.Lines("testdata/tests.txt")
	for i, line := range lines {
		if output := Part1(line); output != expected[i] {
			t.Errorf("Test %d expected %d, got %d", i+1, expected[i], output)
		}
	}
}
