package day06

import (
	_ "embed"
	"errors"
	"strings"
)

//go:embed input.txt
var content string

func tuningTables(input string, part int) (ans int, err error) {
	var step int
	if part == 1 {
		step = 4
	} else {
		step = 14
	}
	for _, line := range strings.Split(input, "\n") {
		if len(line) < 4 {
			return ans, errors.New("invalid input")
		}
		// freq holds the frequency of each letter in step consecutive letters
		freq := make([]int, 26)
		for i := 0; i < step; i++ {
			freq[line[i]-'a']++
		}
		for i := step; i < len(line); i++ {
			// Check if current window has duplicates
			if !hasDuplicates(freq) {
				return i, nil
			}
			freq[line[i-step]-'a']--
			freq[line[i]-'a']++
		}
	}
	return ans, nil
}

func hasDuplicates(freq []int) bool {
	for _, f := range freq {
		if f > 1 {
			return true
		}
	}
	return false
}
