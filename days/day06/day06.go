// Package day06 provides solutions for day 6 of AoC
package day06

import (
	"fmt"
	"strconv"

	"github.com/dlclark/regexp2"
	"github.com/imaspen/aoc-2015-go/inputreader"
)

type lights = [1000][1000]int

type operator int

const (
	turnOn operator = iota + 1
	turnOff
	toggle
)

type instruction struct {
	Operator       operator
	X1, Y1, X2, Y2 int
}

// Run runs the specified part, returning true on success or false if the part is unrecognized
func Run(part int) bool {
	switch part {
	case 1:
		fmt.Println(Part1(inputreader.Lines("data/06.txt")))
		break
	case 2:
		fmt.Println(Part2(inputreader.Lines("data/06.txt")))
		break
	default:
		return false
	}

	return true
}

func parseOperatorString(str string) operator {
	switch str {
	case "turn on":
		return turnOn
	case "turn off":
		return turnOff
	case "toggle":
		return toggle
	default:
		panic(fmt.Sprintf("Unrecognized operator: %s", str))
	}
}

func parseNumber(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf(`Error parsing "%s" as number: %s`, str, err.Error()))
	}
	return i
}

var parser *regexp2.Regexp = regexp2.MustCompile(
	`^(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)$`,
	regexp2.None,
)

func parseInstructionString(string string) instruction {
	match, err := parser.FindStringMatch(string)
	if err != nil {
		panic(fmt.Sprintf("Error parsing instruction: %s", err.Error()))
	}
	groups := match.Groups()
	return instruction{
		parseOperatorString(groups[1].String()),
		parseNumber(groups[2].String()),
		parseNumber(groups[3].String()),
		parseNumber(groups[4].String()),
		parseNumber(groups[5].String()),
	}
}

func parseInstructionStrings(strings []string) []instruction {
	instructions := make([]instruction, 0, len(strings))
	for _, line := range strings {
		instructions = append(instructions, parseInstructionString(line))
	}
	return instructions
}

func executeOperatorOnCell(operator operator, cell *int) {
	switch operator {
	case turnOn:
		*cell = 1
		break
	case turnOff:
		*cell = 0
		break
	case toggle:
		if *cell != 0 {
			*cell = 0
		} else {
			*cell = 1
		}
		break
	default:
		panic(fmt.Sprintf("Unrecognized operator: %d", operator))
	}
}

func executeInstructionOnLights(instruction instruction, lights *lights) {
	for y := instruction.Y1; y <= instruction.Y2; y++ {
		row := &lights[y]
		for x := instruction.X1; x <= instruction.X2; x++ {
			executeOperatorOnCell(instruction.Operator, &row[x])
		}
	}
}

func executeOperatorOnBrightnessCell(operator operator, cell *int) {
	switch operator {
	case turnOn:
		*cell += 1
		break
	case turnOff:
		*cell -= 1
		if *cell < 0 {
			*cell = 0
		}
		break
	case toggle:
		*cell += 2
		break
	default:
		panic(fmt.Sprintf("Unrecognized operator: %d", operator))
	}
}

func executeInstructionOnBrightnessLights(instruction instruction, lights *lights) {
	for y := instruction.Y1; y <= instruction.Y2; y++ {
		row := &lights[y]
		for x := instruction.X1; x <= instruction.X2; x++ {
			executeOperatorOnBrightnessCell(instruction.Operator, &row[x])
		}
	}
}

func countActiveLights(lights lights) int {
	count := 0

	for _, row := range lights {
		for _, cell := range row {
			count += cell
		}
	}

	return count
}

func Part1(lines []string) int {
	lights := lights{}
	instructions := parseInstructionStrings(lines)

	for _, instruction := range instructions {
		executeInstructionOnLights(instruction, &lights)
	}

	return countActiveLights(lights)
}

func Part2(lines []string) int {
	lights := lights{}
	instructions := parseInstructionStrings(lines)

	for _, instruction := range instructions {
		executeInstructionOnBrightnessLights(instruction, &lights)
	}

	return countActiveLights(lights)
}
