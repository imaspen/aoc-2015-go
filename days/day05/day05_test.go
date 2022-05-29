package day05

import (
	"testing"

	"github.com/imaspen/aoc-2015-go/inputreader"
)

func TestPart1(t *testing.T) {
	lines := inputreader.Lines("testdata/part1tests.txt")
	if output := Part1(lines); output != 2 {
		t.Errorf("Part 1 expected 2, got %d", output)
	}
}

func TestPart2(t *testing.T) {
	lines := inputreader.Lines("testdata/part2tests.txt")
	if output := Part2(lines); output != 2 {
		t.Errorf("Part 2 expected 2, got %d", output)
	}
}
