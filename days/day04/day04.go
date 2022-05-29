// Package day04 provides solutions for day 4 of AoC
package day04

import (
	"crypto/md5"
	"fmt"
	"strconv"

	"github.com/imaspen/aoc-2015-go/inputreader"
)

// Run runs the specified part, returning true on success or false if the part is unrecognized
func Run(part int) bool {
	switch part {
	case 1:
		fmt.Println(Part1(inputreader.FirstString("data/04.txt")))
		break
	case 2:
		fmt.Println(Part2(inputreader.FirstString("data/04.txt")))
		break
	default:
		return false
	}

	return true
}

// checkHashForNZeroes checks the provided hash that the first N nibbles are 0
func checkHashForNZeroes(hash [16]byte, n int) bool {
	for i := 0; i < n; i += 2 {
		if i+1 >= n {
			if hash[i/2] > 15 {
				return false
			}
		} else {
			if hash[i/2] != 0 {
				return false
			}
		}
	}

	return true
}

// calculateInputForNZeroes gets the number to append to prefix to generate an MD5 hash with N leading 0 nibbles
func calculateInputForNZeroes(prefix string, n int) int {
	i := 0

	for {
		hash := md5.Sum([]byte(prefix + strconv.Itoa(i)))

		if checkHashForNZeroes(hash, n) {
			return i
		}

		i++
	}
}

func Part1(line string) int {
	return calculateInputForNZeroes(line, 5)
}

func Part2(line string) int {
	return calculateInputForNZeroes(line, 6)
}
