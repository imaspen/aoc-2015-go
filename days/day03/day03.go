// Package day03 provides solutions for day 3 of AoC
package day03

import (
	"fmt"
	"strings"

	"github.com/imaspen/aoc-2015-go/inputreader"
)

// Run runs the specified part, returning true on success or false if the part is unrecognized
func Run(part int) bool {
	switch part {
	case 1:
		fmt.Println(Part1(inputreader.FirstString("data/03.txt")))
		break
	case 2:
		fmt.Println(Part2(inputreader.FirstString("data/03.txt")))
		break
	default:
		return false
	}

	return true
}

func Part1(line string) int {
	visited := map[int]map[int]struct{}{0: {0: {}}}
	chars := strings.Split(line, "")
	x, y := 0, 0
	for _, char := range chars {
		switch char {
		case ">":
			x++
			break
		case "<":
			x--
			break
		case "^":
			y++
			break
		case "v":
			y--
			break
		default:
			panic(fmt.Sprintf("unrecognized direction char: %s", char))
		}

		if visited[y] == nil {
			visited[y] = map[int]struct{}{}
		}
		visited[y][x] = struct{}{}
	}

	total := 0
	for _, row := range visited {
		total += len(row)
	}
	return total
}

func Part2(line string) int {
	visited := map[int]map[int]struct{}{0: {0: {}}}
	chars := strings.Split(line, "")
	x1, y1, x2, y2 := 0, 0, 0, 0
	for i, char := range chars {
		switch char {
		case ">":
			if i%2 == 0 {
				x1++
			} else {
				x2++
			}
			break
		case "<":
			if i%2 == 0 {
				x1--
			} else {
				x2--
			}
			break
		case "^":
			if i%2 == 0 {
				y1++
			} else {
				y2++
			}
			break
		case "v":
			if i%2 == 0 {
				y1--
			} else {
				y2--
			}
			break
		default:
			panic(fmt.Sprintf("unrecognized direction char: %s", char))
		}

		if i%2 == 0 {
			if visited[y1] == nil {
				visited[y1] = map[int]struct{}{}
			}
			visited[y1][x1] = struct{}{}
		} else {
			if visited[y2] == nil {
				visited[y2] = map[int]struct{}{}
			}
			visited[y2][x2] = struct{}{}
		}
	}

	total := 0
	for _, row := range visited {
		total += len(row)
	}
	return total
}
