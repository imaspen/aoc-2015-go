package day06

import (
	"testing"

	"github.com/imaspen/aoc-2015-go/inputreader"
)

func TestPart1(t *testing.T) {
	expected := [...]int{1000 * 1000, 1000, 0}
	lines := inputreader.Lines("testdata/part1tests.txt")
	for i, line := range lines {
		if output := Part1([]string{line}); output != expected[i] {
			t.Errorf("Part 1 test %d expected %d, got %d", i+1, expected[i], output)
		}
	}
	totalExpected := (1000 * 1000) - 1000 - (2 * 2)
	if output := Part1(lines); output != totalExpected {
		t.Errorf("Part 1 combined test expected %d, got %d", totalExpected, output)
	}
}

func TestPart2(t *testing.T) {
	expected := [...]int{4, 0, 2 * 1000 * 1000}
	lines := inputreader.Lines("testdata/part2tests.txt")
	for i, line := range lines {
		if output := Part2([]string{line}); output != expected[i] {
			t.Errorf("Part 2 test %d expected %d, got %d", i+1, expected[i], output)
		}
	}
	totalExpected := 2*1000*1000 + 2
	if output := Part2(lines); output != totalExpected {
		t.Errorf("Part 1 combined test expected %d, got %d", totalExpected, output)
	}
}
