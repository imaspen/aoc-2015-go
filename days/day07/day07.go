// Package day07 provides solutions for day 7 of AoC
package day07

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
	"github.com/imaspen/aoc-2015-go/inputreader"
)

// Run runs the specified part, returning true on success or false if the part is unrecognized
func Run(part int) bool {
	switch part {
	case 1:
		fmt.Println(Part1(inputreader.Lines("data/07.txt"))["a"])
		break
	case 2:
		fmt.Println(Part2(inputreader.Lines("data/07.txt"))["a"])
		break
	default:
		return false
	}

	return true
}

type Signals map[string]uint16

type operator int

const (
	set operator = iota + 1
	and
	or
	lshift
	rshift
	not
)

type instruction struct {
	Operator               operator
	Input1, Input2, Target string
	Value                  int
}

var setRegexp regexp2.Regexp = *regexp2.MustCompile(`^([a-z]+|\d+)$`, regexp2.None)
var andOrRegexp regexp2.Regexp = *regexp2.MustCompile(`^([a-z]+|\d+) (AND|OR) ([a-z]+)$`, regexp2.None)
var shiftRegexp regexp2.Regexp = *regexp2.MustCompile(`^([a-z]+) ([LR]SHIFT) (\d+)$`, regexp2.None)
var notRegexp regexp2.Regexp = *regexp2.MustCompile(`^NOT ([a-z]+)$`, regexp2.None)

func getInstructionFromString(s string) instruction {
	parts := strings.Split(s, " -> ")
	target := parts[1]

	setRes, setErr := setRegexp.FindStringMatch(parts[0])
	if setErr != nil {
		panic(fmt.Sprintf("Error checking for set: %s", setErr.Error()))
	}
	if setRes != nil {
		str := setRes.Groups()[1].String()
		value, err := strconv.Atoi(str)
		if err == nil {
			return instruction{
				set,
				"",
				"",
				target,
				value,
			}
		} else {
			return instruction{
				set,
				str,
				"",
				target,
				0,
			}
		}
	}

	andOrRes, andOrErr := andOrRegexp.FindStringMatch(parts[0])
	if andOrErr != nil {
		panic(fmt.Sprintf("Error checking for and/or: %s", andOrErr.Error()))
	}
	if andOrRes != nil {
		groups := andOrRes.Groups()

		var op operator
		switch groups[2].String() {
		case "AND":
			op = and
			break
		case "OR":
			op = or
			break
		default:
			panic(fmt.Sprintf("Unrecognized and/or operator: %s", groups[2].String()))
		}

		value, err := strconv.Atoi(andOrRes.Groups()[1].String())
		if err == nil {
			return instruction{
				op,
				"",
				groups[3].String(),
				target,
				value,
			}
		} else {
			return instruction{
				op,
				groups[1].String(),
				groups[3].String(),
				target,
				0,
			}
		}
	}

	shiftRes, shiftErr := shiftRegexp.FindStringMatch(parts[0])
	if shiftErr != nil {
		panic(fmt.Sprintf("Error checking for shift: %s", shiftErr.Error()))
	}
	if shiftRes != nil {
		groups := shiftRes.Groups()

		var op operator
		switch groups[2].String() {
		case "LSHIFT":
			op = lshift
			break
		case "RSHIFT":
			op = rshift
			break
		default:
			panic(fmt.Sprintf("Unrecognized shift operator: %s", groups[2].String()))
		}

		value, err := strconv.Atoi(shiftRes.Groups()[3].String())
		if err != nil {
			panic(fmt.Sprintf("Error parsing shift value: %s", err.Error()))
		}

		return instruction{
			op,
			groups[1].String(),
			"",
			target,
			value,
		}
	}

	notRes, notErr := notRegexp.FindStringMatch(parts[0])
	if notErr != nil {
		panic(fmt.Sprintf("Error checking for not: %s", notErr.Error()))
	}
	if notRes != nil {
		return instruction{
			not,
			notRes.Groups()[1].String(),
			"",
			target,
			0,
		}
	}

	panic(fmt.Sprintf(`Failed to parse instruction "%s"`, s))
}

func applyInstructionToSignals(instruction instruction, signals Signals) bool {
	switch instruction.Operator {
	case set:
		if len(instruction.Input1) > 0 {
			val, exists := signals[instruction.Input1]
			if exists {
				signals[instruction.Target] = val
			} else {
				return false
			}
		} else {
			_, exists := signals[instruction.Target]
			if !exists {
				signals[instruction.Target] = uint16(instruction.Value)
			}
		}
		return true
	case and:
		val2, exists2 := signals[instruction.Input2]
		if !exists2 {
			return false
		}
		if len(instruction.Input1) > 0 {
			val1, exists1 := signals[instruction.Input1]
			if !exists1 {
				return false
			}
			signals[instruction.Target] = val1 & val2
		} else {
			signals[instruction.Target] = uint16(instruction.Value) & val2
		}
		return true
	case or:
		val2, exists2 := signals[instruction.Input2]
		if !exists2 {
			return false
		}
		if len(instruction.Input1) > 0 {
			val1, exists1 := signals[instruction.Input1]
			if !exists1 {
				return false
			}
			signals[instruction.Target] = val1 | val2
		} else {
			signals[instruction.Target] = uint16(instruction.Value) | val2
		}
		return true
	case lshift:
		val, exists := signals[instruction.Input1]
		if !exists {
			return false
		}
		signals[instruction.Target] = val << instruction.Value
		return true
	case rshift:
		val, exists := signals[instruction.Input1]
		if !exists {
			return false
		}
		signals[instruction.Target] = val >> instruction.Value
		return true
	case not:
		val, exists := signals[instruction.Input1]
		if !exists {
			return false
		}
		signals[instruction.Target] = val ^ math.MaxUint16
		return true
	default:
		panic(fmt.Sprintf("Unrecognized operator: %d", instruction.Operator))
	}
}

func Part1(input []string) Signals {
	signals := Signals{}
	var instructions []instruction

	for _, line := range input {
		instructions = append(instructions, getInstructionFromString(line))
	}

	for len(instructions) > 0 {
		for i, instruction := range instructions {
			if applyInstructionToSignals(instruction, signals) {
				instructions = append(instructions[:i], instructions[i+1:]...)
				break
			}
		}
	}

	return signals
}

func Part2(input []string) Signals {
	override := Part1(input)["a"]

	signals := Signals{"b": override}
	var instructions []instruction

	for _, line := range input {
		instructions = append(instructions, getInstructionFromString(line))
	}

	for len(instructions) > 0 {
		for i, instruction := range instructions {
			if applyInstructionToSignals(instruction, signals) {
				instructions = append(instructions[:i], instructions[i+1:]...)
				break
			}
		}
	}

	return signals
}
