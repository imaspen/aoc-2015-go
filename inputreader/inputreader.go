// Package inputreader handles the common types of input files used for AoC
package inputreader

import (
	"bufio"
	"os"
)

// FirstString gets the first line of the passed file path as a string.
func FirstString(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return ""
	}
	return scanner.Text()
}

// Lines gets the entire file passed via file path as an array of strings.
// One string per line in the file.
func Lines(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var strings []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}

	return strings
}
