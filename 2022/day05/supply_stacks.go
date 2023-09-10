package day05

import (
	"aoc/datastructures/deque"
	_ "embed"
	"errors"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var content string

func supplyStacks(input string, part int) (string, error) {
	var ans string

	// The input is divided into 2 parts: starting deques of crates and crates rearrangement
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return ans, errors.New("invalid empty input")
	}

	// Initialize starting deques
	var deques []*deque.Deque[string]
	for i := 0; i < len(lines[0]); i += 4 {
		deques = append(deques, deque.New[string]())
	}

	// 1. Populate starting deques of crates
	index, err := buildCrateDeques(lines, deques)
	if err != nil {
		return ans, err
	}

	// 2. Rearrange deque of crates
	err = rearrangeCrateDeques(index, lines, deques, part)
	if err != nil {
		return ans, err
	}

	for _, d := range deques {
		res, err := d.Back()
		if err != nil { // empty deque, append space to ans instead
			ans += " "
		} else {
			ans += res
		}
	}
	return ans, nil
}

// buildCrateDeques returns the first index of line for crates rearrangement
func buildCrateDeques(lines []string, deques []*deque.Deque[string]) (int, error) {
	for i, line := range lines {
		if line == "" { // Start of crates rearrangement
			return i + 1, nil
		}
		// Iterate over 4 columns at a time and look for '['
		// j is index of line, while k is index of deque
		for j, k := 0, 0; j < len(line) && k < len(deques); j, k = j+4, k+1 {
			// If there's crate, append to deque
			if line[j] == '[' {
				deques[k].PushFront(string(line[j+1]))
			}
		}
	}
	return 0, errors.New("invalid input: no crates rearrangement found")
}

// rearrangeCrateDeques rearrange deques of crates given a list of procedure
func rearrangeCrateDeques(index int, lines []string, deques []*deque.Deque[string], part int) error {
	for i := index; i < len(lines); i++ {
		// A procedure follows this format 'move x from y to z'
		// x: The number of crates to move
		// y: The origin deque from which crates are removed
		// z: The destination deque to which crates are added
		subs := strings.Split(lines[i], " ")
		if len(subs) != 6 {
			return errors.New("invalid procedure input")
		}

		n, err := strconv.Atoi(subs[1])
		if err != nil {
			return err
		}
		s1, err := strconv.Atoi(subs[3])
		if err != nil {
			return err
		}
		s2, err := strconv.Atoi(subs[5])
		if err != nil {
			return err
		}

		// Retain the order of crates to move
		var movedCrates []string
		for j := 0; j < n; j++ {
			res, err := deques[s1-1].PopBack()
			// If deque is already empty, then no-op
			if err != nil {
				continue
			}
			movedCrates = append(movedCrates, res)
		}

		if part == 2 {
			slices.Reverse(movedCrates)
		}
		for _, crate := range movedCrates {
			deques[s2-1].PushBack(crate)
		}
	}
	return nil
}
