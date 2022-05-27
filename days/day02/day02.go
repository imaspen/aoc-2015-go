// Package day02 provides solutions for day 2 of AoC
package day02

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/imaspen/aoc-2015-go/inputreader"
)

// Run runs the specified part, returning true on success or false if the part is unrecognized
func Run(part int) bool {
	switch part {
	case 1:
		fmt.Println(Part1(inputreader.Lines("data/02.txt")))
		break
	case 2:
		fmt.Println(Part2(inputreader.Lines("data/02.txt")))
		break
	default:
		return false
	}

	return true
}

func parseWrappingPaperDimensions(line string) (l, w, h int) {
	parts := strings.Split(line, "x")
	if len(parts) != 3 {
		panic(fmt.Sprintf("Unexpected wrapping paper dimensions length: %s", line))
	}

	l, lerr := strconv.Atoi(parts[0])
	w, werr := strconv.Atoi(parts[1])
	h, herr := strconv.Atoi(parts[2])
	if lerr != nil || werr != nil || herr != nil {
		panic(fmt.Sprintf("Unexpected wrapping paper dimensions parameters: %s", line))
	}

	return l, w, h
}

// calculateWrappingPaperArea using 2*l*w + 2*w*h + 2*h*l + (area of smallest side)
func calculateWrappingPaperArea(l, w, h int) int {
	edges := [...]int{
		l * w,
		w * h,
		h * l,
	}

	min := math.MaxInt
	for _, edge := range edges {
		if edge < min {
			min = edge
		}
	}

	return 2*(edges[0]+edges[1]+edges[2]) + min
}

func Part1(lines []string) int {
	total := 0
	for _, line := range lines {
		total += calculateWrappingPaperArea(parseWrappingPaperDimensions(line))
	}
	return total
}

func calculateRibbonLength(l, w, h int) int {
	max := 0
	for _, edge := range [3]int{l, w, h} {
		if edge > max {
			max = edge
		}
	}

	return 2*(l+w+h) - 2*max + l*w*h
}

func Part2(lines []string) int {
	total := 0
	for _, line := range lines {
		total += calculateRibbonLength(parseWrappingPaperDimensions(line))
	}
	return total
}
