package day01

import (
	"aoc/util"
	_ "embed"
	"fmt"
	"strings"
	"unicode"
)

//go:embed input.txt
var content string // Embed file content into string

func init() {
	// Preprocess input here, so it applies to test
	// Append a newline to content for consistency
	content += "\n"
}

func trebutchetPart1(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		firstDigit, secondDigit := getFirstDigit(line, 1), getSecondDigit(line, 1)
		sum += firstDigit*10 + secondDigit
	}
	return sum
}

func trebutchetPart2(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		firstDigit, secondDigit := getFirstDigit(line, 2), getSecondDigit(line, 2)
		sum += firstDigit*10 + secondDigit
	}
	return sum
}

func getFirstDigit(s string, part int) int {
	firstDigit, minDigitIndex := 0, len(s)
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			if i < minDigitIndex {
				firstDigit, minDigitIndex = int(s[i]-'0'), i
			}
		}
	}
	if part == 2 {
		// Do a scan of i-letters strings, possible i = 3, 4, 5
		for i := 3; i <= 5; i++ {
			for j := 0; j < len(s)-i+1; j++ {
				if digit, err := digitFromStr(s[j : j+i]); err == nil {
					if j < minDigitIndex {
						firstDigit, minDigitIndex = *digit, j
					}
				}
			}
		}
	}
	return firstDigit
}

func getSecondDigit(s string, part int) int {
	s = util.ReverseString(s)
	secondDigit, minDigitIndex := 0, len(s)
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			if i < minDigitIndex {
				secondDigit, minDigitIndex = int(s[i]-'0'), i
			}
		}
	}
	if part == 2 {
		// Do a scan of i-letters strings, possible i = 3, 4, 5
		for i := 3; i <= 5; i++ {
			for j := 0; j < len(s)-i+1; j++ {
				if digit, err := digitFromStr(util.ReverseString(s[j : j+i])); err == nil {
					if j < minDigitIndex {
						secondDigit, minDigitIndex = *digit, j
					}
				}
			}
		}
	}
	return secondDigit
}

// 3, 4, 5
func digitFromStr(s string) (*int, error) {
	switch s {
	case "one":
		return util.IntPtr(1), nil
	case "two":
		return util.IntPtr(2), nil
	case "three":
		return util.IntPtr(3), nil
	case "four":
		return util.IntPtr(4), nil
	case "five":
		return util.IntPtr(5), nil
	case "six":
		return util.IntPtr(6), nil
	case "seven":
		return util.IntPtr(7), nil
	case "eight":
		return util.IntPtr(8), nil
	case "nine":
		return util.IntPtr(9), nil
	default:
		return nil, fmt.Errorf("invalid digit string: %s", s)
	}
}
