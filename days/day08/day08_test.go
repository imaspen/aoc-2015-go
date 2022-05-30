package day08

import (
	"testing"

	"github.com/imaspen/aoc-2015-go/inputreader"
)

func TestPart1(t *testing.T) {
	expected := [...]int{2, 2, 3, 5}
	lines := inputreader.Lines("testdata/tests.txt")
	for i, line := range lines {
		if output := Part1([]string{line}); output != expected[i] {
			t.Errorf("Part 1 test %d expected %d, got %d", i+1, expected[i], output)
		}
	}

	if output := Part1(lines); output != 12 {
		t.Errorf("Part 1 combined test expected 12, got %d", output)
	}
}

func TestPart2(t *testing.T) {
	expected := [...]int{4, 4, 6, 5}
	lines := inputreader.Lines("testdata/tests.txt")
	for i, line := range lines {
		if output := Part2([]string{line}); output != expected[i] {
			t.Errorf("Part 2 test %d expected %d, got %d", i+1, expected[i], output)
		}
	}

	if output := Part2(lines); output != 19 {
		t.Errorf("Part 2 combined test expected 19, got %d", output)
	}
}
