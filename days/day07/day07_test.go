package day07

import (
	"testing"

	"github.com/imaspen/aoc-2015-go/inputreader"
)

func TestPart1(t *testing.T) {
	expected := Signals{
		"d": 72,
		"e": 507,
		"f": 492,
		"g": 114,
		"h": 65412,
		"i": 65079,
		"x": 123,
		"y": 456,
	}
	lines := inputreader.Lines("testdata/part1tests.txt")
	output := Part1(lines)
	for k, v := range output {
		if expected[k] != v {
			t.Errorf(`Signal "%s" was expected to be %d, instead was %d.`, k, expected[k], v)
		}
	}
	for k, v := range expected {
		if output[k] != v {
			t.Errorf(`Signal "%s" was expected to be %d, instead was %d.`, k, v, output[k])
		}
	}
}
