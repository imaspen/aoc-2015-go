package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/imaspen/aoc-2015-go/days/day01"
	"github.com/imaspen/aoc-2015-go/days/day02"
	"github.com/imaspen/aoc-2015-go/days/day03"
)

// exit prints usage information and then exits with code 1.
// Used to provide generic information for exceptional inputs.
func exit() {
	fmt.Println("Usage: aoc {day:1..25} {part:1|2}")
	os.Exit(1)
}

func main() {
	if len(os.Args) != 3 {
		exit()
	}

	day, dayError := strconv.Atoi(os.Args[1])
	if dayError != nil {
		exit()
	}

	part, partError := strconv.Atoi(os.Args[2])
	if partError != nil {
		exit()
	}

	result := false

	switch day {
	case 1:
		result = day01.Run(part)
		break
	case 2:
		result = day02.Run(part)
		break
	case 3:
		result = day03.Run(part)
		break
	default:
		fmt.Printf("Unrecognized day: %d.\n", day)
		exit()
	}

	if !result {
		fmt.Printf("Unrecognized part: %d.\n", part)
		exit()
	}
}
