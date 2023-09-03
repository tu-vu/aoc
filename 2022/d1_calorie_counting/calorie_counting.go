package d1_calorie_counting

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func calorieCountingPart1(filePath string) int {
	// Max calo carried by an elf
	maxCalo := -1

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return maxCalo
	}
	defer file.Close()

	// Parse the data
	caloSoFar := 0 // Total calo carried by an elf so far
	reader := bufio.NewReader(file)
	for {
		caloStr, err := reader.ReadString('\n')
		if err != nil {
			// EOF
			break
		}

		// Trim trailing new line character
		caloStr = strings.TrimSuffix(caloStr, "\n")

		caloInt, err := strconv.Atoi(caloStr)
		if err != nil {
			// End of inventory for current elf
			// Check if total calo exceeds max before reset
			if caloSoFar > maxCalo {
				maxCalo = caloSoFar
			}
			caloSoFar = 0
		} else {
			caloSoFar += caloInt
		}
	}
	return maxCalo
}

func calorieCountingPart2(filePath string) int {
	maxCalo1 := -1 // Top 1 calo carried by an elf
	maxCalo2 := -1 // Top 2 calo carried by an elf
	maxCalo3 := -1 // Top 3 calo carried by an elf

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return -1
	}
	defer file.Close()

	// Parse the data
	caloSoFar := 0 // Total calo carried by an elf so far
	reader := bufio.NewReader(file)
	for {
		caloStr, err := reader.ReadString('\n')
		if err != nil {
			// EOF
			break
		}
		// Trim trailing new line character
		caloStr = strings.TrimSuffix(caloStr, "\n")

		caloInt, err := strconv.Atoi(caloStr)
		if err != nil {
			// End of inventory for current elf
			// Check if total calo exceeds each max, swap their values if so
			// This is because we want to retain the values of each max as
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
			caloSoFar += caloInt
		}
	}
	return maxCalo1 + maxCalo2 + maxCalo3
}
