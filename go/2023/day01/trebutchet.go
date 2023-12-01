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
		firstDigit, secondDigit := 0, 0
		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				firstDigit = int(line[i] - '0')
				break
			}
		}
		for j := len(line) - 1; j >= 0; j-- {
			if unicode.IsDigit(rune(line[j])) {
				secondDigit = int(line[j] - '0')
				break
			}
		}
		sum += firstDigit*10 + secondDigit
	}
	return sum
}

func trebutchetPart2(input string) int {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		firstDigit, secondDigit := getFirstDigit(line), getSecondDigit(util.ReverseString(line))
		fmt.Printf("%d%d\n", firstDigit, secondDigit)
		sum += firstDigit*10 + secondDigit
	}
	return sum
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

func getFirstDigit(s string) int {
	firstDigit, minDigitIndex := 0, len(s)
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			if i < minDigitIndex {
				firstDigit, minDigitIndex = int(s[i]-'0'), i
			}
		}
	}
	// Do a scan of 3-letters strings
	if len(s) > 2 {
		for i := 0; i < len(s)-2; i++ {
			if digit, err := digitFromStr(s[i : i+3]); err == nil {
				if i < minDigitIndex {
					firstDigit, minDigitIndex = *digit, i
				}
			}
		}
	}
	// Do a scan of 4-letters strings
	if len(s) > 3 {
		for i := 0; i < len(s)-3; i++ {
			if digit, err := digitFromStr(s[i : i+4]); err == nil {
				if i < minDigitIndex {
					firstDigit, minDigitIndex = *digit, i
				}
			}
		}
	}
	// Do a scan of 5-letters strings
	if len(s) > 4 {
		for i := 0; i < len(s)-4; i++ {
			if digit, err := digitFromStr(s[i : i+5]); err == nil {
				if i < minDigitIndex {
					firstDigit, minDigitIndex = *digit, i
				}
			}
		}
	}
	return firstDigit
}

func getSecondDigit(s string) int {
	secondDigit, minDigitIndex := 0, len(s)
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			if i < minDigitIndex {
				secondDigit, minDigitIndex = int(s[i]-'0'), i
			}
		}
	}
	// Do a scan of 3-letters strings
	if len(s) > 2 {
		for i := 0; i < len(s)-2; i++ {
			if digit, err := digitFromStr(util.ReverseString(s[i : i+3])); err == nil {
				if i < minDigitIndex {
					secondDigit, minDigitIndex = *digit, i
				}
			}
		}
	}
	// Do a scan of 4-letters strings
	if len(s) > 3 {
		for i := 0; i < len(s)-3; i++ {
			if digit, err := digitFromStr(util.ReverseString(s[i : i+4])); err == nil {
				if i < minDigitIndex {
					secondDigit, minDigitIndex = *digit, i
				}
			}
		}
	}
	// Do a scan of 5-letters strings
	if len(s) > 4 {
		for i := 0; i < len(s)-4; i++ {
			if digit, err := digitFromStr(util.ReverseString(s[i : i+5])); err == nil {
				if i < minDigitIndex {
					secondDigit, minDigitIndex = *digit, i
				}
			}
		}
	}
	return secondDigit
}
