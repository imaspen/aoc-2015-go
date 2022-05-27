// Package day01 provides solutions for day 1 of AoC
package day01

import (
	"fmt"
	"strings"

	"github.com/imaspen/aoc-2015-go/inputreader"
)

// Run runs the specified part, returning true on success or false if the part is unrecognized
func Run(part int) bool {
	switch part {
	case 1:
		fmt.Println(Part1(inputreader.FirstString("data/01.txt")))
		break
	case 2:
		fmt.Println(Part2(inputreader.FirstString("data/01.txt")))
		break
	default:
		return false
	}

	return true
}

func Part1(input string) int {
	chars := strings.Split(input, "")

	count := 0
	for _, char := range chars {
		if char == "(" {
			count++
		} else {
			count--
		}
	}

	return count
}

func Part2(input string) int {
	chars := strings.Split(input, "")

	count := 0
	for i, char := range chars {
		if char == "(" {
			count++
		} else {
			count--
		}
		if count == -1 {
			return i + 1
		}
	}
	return 0
}
