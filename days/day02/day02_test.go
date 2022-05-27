package day02

import (
	"testing"

	"github.com/imaspen/aoc-2015-go/inputreader"
)

func TestParseWrappingPaperDimensions(t *testing.T) {
	expected := [...][3]int{{2, 3, 4}, {1, 1, 10}}
	lines := inputreader.Lines("testdata/tests.txt")
	for i, line := range lines {
		e := expected[i]
		if l, w, h := parseWrappingPaperDimensions(line); l != e[0] || w != e[1] || h != e[2] {
			t.Errorf("Test %d expected %dx%dx%d, got %dx%dx%d", i+1, e[0], e[1], e[2], l, w, h)
		}
	}
}

func TestCalculateWrappingPaperArea(t *testing.T) {
	expected := [...]int{58, 43}
	lines := inputreader.Lines("testdata/tests.txt")
	for i, line := range lines {
		if output := calculateWrappingPaperArea(parseWrappingPaperDimensions(line)); output != expected[i] {
			t.Errorf("Test %d expected %d, got %d", i+1, expected[i], output)
		}
	}
}

func TestPart1(t *testing.T) {
	lines := inputreader.Lines("testdata/tests.txt")
	if output := Part1(lines); output != 101 {
		t.Errorf("Part 1 expected 101, got %d", output)
	}
}

func TestCalculateRibbonLength(t *testing.T) {
	expected := [...]int{34, 14}
	lines := inputreader.Lines("testdata/tests.txt")
	for i, line := range lines {
		if output := calculateRibbonLength(parseWrappingPaperDimensions(line)); output != expected[i] {
			t.Errorf("Test %d expected %d, got %d", i+1, expected[i], output)
		}
	}
}

func TestPart2(t *testing.T) {
	lines := inputreader.Lines("testdata/tests.txt")
	if output := Part2(lines); output != 48 {
		t.Errorf("Part 2 expected 48, got %d", output)
	}
}
