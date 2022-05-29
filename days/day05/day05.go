// Package day05 provides solutions for day 5 of AoC
package day05

import (
	"fmt"

	"github.com/dlclark/regexp2"
	"github.com/imaspen/aoc-2015-go/inputreader"
)

type test struct {
	regexp      *regexp2.Regexp
	shouldMatch bool
}

// Run runs the specified part, returning true on success or false if the part is unrecognized
func Run(part int) bool {
	switch part {
	case 1:
		fmt.Println(Part1(inputreader.Lines("data/05.txt")))
		break
	case 2:
		fmt.Println(Part2(inputreader.Lines("data/05.txt")))
		break
	default:
		return false
	}

	return true
}

func doesLinePassTests(line string, tests []test) bool {
	for _, test := range tests {
		match, error := test.regexp.MatchString(line)
		if error != nil {
			panic(fmt.Sprintf("Regexp timed out"))
		}
		if match != test.shouldMatch {
			return false
		}
	}

	return true
}

func countOfLinesPassingTests(lines []string, tests []test) int {
	count := 0

	for _, line := range lines {
		if doesLinePassTests(line, tests) {
			count++
		}
	}

	return count
}

func Part1(lines []string) int {
	containsThreeVowels := regexp2.MustCompile(`[aeiou].*[aeiou].*[aeiou]`, regexp2.None)
	containsRepeatedLetter := regexp2.MustCompile(`([a-z])\1+`, regexp2.None)
	containsNoPredefinedString := regexp2.MustCompile(`(ab|cd|pq|xy)`, regexp2.None)

	tests := []test{
		{containsThreeVowels, true},
		{containsRepeatedLetter, true},
		{containsNoPredefinedString, false},
	}

	return countOfLinesPassingTests(lines, tests)
}

func Part2(lines []string) int {
	containsRepeatedTwoLetterString := regexp2.MustCompile(`([a-z]{2}).*\1`, regexp2.None)
	containsRepeatedLetterWithOneGap := regexp2.MustCompile(`([a-z]).\1`, regexp2.None)

	tests := []test{
		{containsRepeatedTwoLetterString, true},
		{containsRepeatedLetterWithOneGap, true},
	}

	return countOfLinesPassingTests(lines, tests)
}
