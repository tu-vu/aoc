package day04

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var content string

func campCleanupPart1(input string) (int, error) {
	var ans int
	for _, line := range strings.Split(input, "\n") {
		// assignments[i] is section assignment for elf i
		assignments, err := parseInput(line)
		if err != nil {
			return ans, err
		}
		// Check if there's any fully overlapped assignment between elf 1 and 2
		if isFullyOverlapped(assignments) {
			ans++
		}
	}
	return ans, nil
}

func campCleanupPart2(input string) (int, error) {
	var ans int
	for _, line := range strings.Split(input, "\n") {
		// assignments[i] is section assignment for elf i + 1
		assignments, err := parseInput(line)
		if err != nil {
			return ans, err
		}
		// Check if there's any overlapped assignment between elf 1 and 2
		if isOverlapped(assignments) {
			ans++
		}
	}
	return ans, nil
}

func parseInput(line string) ([][]int, error) {
	// pair[i] is section assignment for elf i + 1
	pair := strings.Split(line, ",")

	// Get start and end section assignment for each elf
	assignments := make([][]int, len(pair))
	for i, p := range pair {
		assignment := make([]int, 2)
		for j, sectionStr := range strings.Split(p, "-") {
			section, err := strconv.Atoi(sectionStr)
			if err != nil {
				return nil, err
			}
			assignment[j] = section
		}
		assignments[i] = assignment
	}
	return assignments, nil
}

func isFullyOverlapped(assignments [][]int) bool {
	return (assignments[0][0] <= assignments[1][0] && assignments[1][1] <= assignments[0][1]) ||
		(assignments[1][0] <= assignments[0][0] && assignments[0][1] <= assignments[1][1])
}

func isOverlapped(assignments [][]int) bool {
	return (assignments[1][0] <= assignments[0][0] && assignments[0][0] <= assignments[1][1]) ||
		(assignments[0][0] <= assignments[1][0] && assignments[1][0] <= assignments[0][1])
}
