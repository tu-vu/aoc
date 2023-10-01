package day06

import (
	_ "embed"
	"errors"
	"strings"
)

//go:embed input.txt
var content string

func tuningTables(input string, part int) (ans int, err error) {
	var n int
	if part == 1 {
		n = 4
	} else {
		n = 14
	}
	for _, line := range strings.Split(input, "\n") {
		if len(line) < n {
			return ans, errors.New("invalid input")
		}
		// freq holds the frequency of each letter in n consecutive letters
		freq := make([]int, 26)
		for i := 0; i < n; i++ {
			freq[line[i]-'a']++
		}
		for i := n; i < len(line); i++ {
			// Check if current window has duplicates
			if !hasDuplicates(freq) {
				return i, nil
			}
			freq[line[i-n]-'a']--
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
