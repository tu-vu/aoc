package day03

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var content string

func rucksackReorganizationPart1(input string) (int, error) {
	var totalPriority int
	for _, line := range strings.Split(input, "\n") {
		lineLength := len(line)
		rucksacks := []string{line[:lineLength/2], line[lineLength/2 : lineLength]}

		// Frequency for each item over group of rucksacks
		// freqInGroup[i][j] maps item i to its frequency in rucksack j
		freqInGroup := getFreqInGroup(rucksacks)

		// Look for item that's found in all rucksacks
		for j := 0; j < len(freqInGroup); j++ {
			if freqInGroup[j][0] > 0 && freqInGroup[j][1] > 0 {
				totalPriority += j + 1
				break
			}
		}
	}
	return totalPriority, nil
}

func rucksackReorganizationPart2(input string) (int, error) {
	var totalPriority int
	// Iterate over 3 lines at a time
	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i += 3 {
		var rucksacks = []string{lines[i], lines[i+1], lines[i+2]}

		// Frequency for each item over group of rucksacks
		// freqInGroup[i][j] maps item i to its frequency in rucksack j
		freqInGroup := getFreqInGroup(rucksacks)

		// Look for item that's found in all rucksacks
		for j := 0; j < len(freqInGroup); j++ {
			if freqInGroup[j][0] > 0 && freqInGroup[j][1] > 0 && freqInGroup[j][2] > 0 {
				totalPriority += j + 1
				break
			}
		}
	}
	return totalPriority, nil
}

func getFreqInGroup(rucksacks []string) map[int]map[int]int {
	freqInGroup := make(map[int]map[int]int)
	for j := 0; j < 52; j++ {
		freqInGroup[j] = make(map[int]int)
	}
	for group, rucksack := range rucksacks {
		for _, item := range rucksack {
			if 'a' <= item && item <= 'z' {
				freqInGroup[int(item-'a')][group]++
			} else {
				freqInGroup[int(item-'A'+26)][group]++
			}
		}
	}
	return freqInGroup
}
