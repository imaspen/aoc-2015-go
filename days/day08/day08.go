// Package day08 provides solutions for day 8 of AoC
package day08

import (
	"fmt"
	"regexp"

	"github.com/imaspen/aoc-2015-go/inputreader"
)

// Run runs the specified part, returning true on success or false if the part is unrecognized
func Run(part int) bool {
	switch part {
	case 1:
		fmt.Println(Part1(inputreader.Lines("data/08.txt")))
		break
	case 2:
		fmt.Println(Part2(inputreader.Lines("data/08.txt")))
		break
	default:
		return false
	}

	return true
}

func Part1(input []string) int {
	regex := regexp.MustCompile(`\\\\|\\"|\\x[0-9a-f]{2}`)
	count := 0
	for _, line := range input {
		parsed := regex.ReplaceAllString(line, ` `)
		count += len(line) - (len(parsed) - 2)
	}
	return count
}

func Part2(input []string) int {
	regex := regexp.MustCompile(`\\|"`)
	count := 0
	for _, line := range input {
		parsed := regex.ReplaceAllString(line, `  `)
		count += (len(parsed) + 2) - len(line)
	}
	return count
}
