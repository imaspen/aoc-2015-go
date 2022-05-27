package day03

import (
	"testing"

	"github.com/imaspen/aoc-2015-go/inputreader"
)

func TestPart1(t *testing.T) {
	expected := [...]int{2, 4, 2}
	lines := inputreader.Lines("testdata/part1tests.txt")
	for i, line := range lines {
		if output := Part1(line); output != expected[i] {
			t.Errorf("Test %d expected %d, got %d", i+1, expected[i], output)
		}
	}
}

func TestPart2(t *testing.T) {
	expected := [...]int{3, 3, 11}
	lines := inputreader.Lines("testdata/part2tests.txt")
	for i, line := range lines {
		if output := Part2(line); output != expected[i] {
			t.Errorf("Test %d expected %d, got %d", i+1, expected[i], output)
		}
	}
}
