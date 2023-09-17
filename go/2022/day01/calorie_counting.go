package day01

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input.txt
var content string // Embed file content into string

func init() {
	// Preprocess input here, so it applies to test
	// Append a newline to content for consistency
	content += "\n"
}

func calorieCountingPart1(input string) int {
	maxCalo := -1  // Max calo carried by an elf
	caloSoFar := 0 // Total calo carried by an elf so far
	for _, line := range strings.Split(input, "\n") {
		calo, err := strconv.Atoi(line)
		if err != nil {
			// End of inventory for current elf
			// Check if total calo exceeds max before reset
			if caloSoFar > maxCalo {
				maxCalo = caloSoFar
			}
			caloSoFar = 0
		} else {
			caloSoFar += calo
		}
	}
	return maxCalo
}

func calorieCountingPart2(input string) int {
	maxCalo1 := -1 // Top 1 calo carried by an elf
	maxCalo2 := -1 // Top 2 calo carried by an elf
	maxCalo3 := -1 // Top 3 calo carried by an elf

	caloSoFar := 0 // Total calo carried by an elf so far
	for _, line := range strings.Split(input, "\n") {
		calo, err := strconv.Atoi(line)
		if err != nil {
			// End of inventory for current elf
			// Check if total calo exceeds each max, swap their values if so
			// This is because we want to persist the values of each max as
			// they are likely to be the candidate for next max
			if caloSoFar > maxCalo1 {
				caloSoFar, maxCalo1 = maxCalo1, caloSoFar
			}
			if caloSoFar > maxCalo2 {
				caloSoFar, maxCalo2 = maxCalo2, caloSoFar
			}
			if caloSoFar > maxCalo3 {
				maxCalo3 = caloSoFar
			}
			// Reset total calo
			caloSoFar = 0
		} else {
			caloSoFar += calo
		}
	}
	return maxCalo1 + maxCalo2 + maxCalo3
}
