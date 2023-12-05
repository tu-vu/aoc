package day04

import (
	"aoc/datastructures/set"
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var content string // Embed file content into string

func scratchCardsPart1(input string) float64 {
	var sum float64
	winningCards, existingCards, err := parseInput(input)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(winningCards); i++ {
		// Populate set of winning numbers
		winningNums := set.New[int]()
		for _, num := range winningCards[i] {
			winningNums.Add(num)
		}
		// Count the number of winning numbers in existing cards
		var count float64
		for j := 0; j < len(existingCards[i]); j++ {
			if winningNums.Has(existingCards[i][j]) {
				count++
			}
		}
		// Update the sum
		if count > 0 {
			sum += math.Pow(2, count-1)
		}
	}
	return sum
}

func scratchCardsPart2(input string) int {
	winningCards, existingCards, err := parseInput(input)
	if err != nil {
		panic(err)
	}
	// Initialize map of card to its number of copies
	sum := len(winningCards)
	cardCopies := make(map[int]int, 0)
	for i := 0; i < len(winningCards); i++ {
		cardCopies[i] = 1
	}
	for i := 0; i < len(winningCards); i++ {
		for cardCopies[i] > 0 {
			// Populate set of winning numbers
			winningNums := set.New[int]()
			for _, num := range winningCards[i] {
				winningNums.Add(num)
			}
			// Count the number of winning numbers in existing cards
			var count int
			for j := 0; j < len(existingCards[i]); j++ {
				if winningNums.Has(existingCards[i][j]) {
					count++
				}
			}
			// Update copies of the later cards based on the count
			for k := i + 1; k < i+1+count; k++ {
				cardCopies[k]++
			}
			// Update the count of copies of card i
			cardCopies[i]--
			// Update the sum
			sum += count
		}
	}
	return sum
}

func parseInput(input string) ([][]int, [][]int, error) {
	winningCards := make([][]int, 0)
	existingCards := make([][]int, 0)
	for _, line := range strings.Split(input, "\n") {
		lineParts := strings.Split(line, ":")
		numsParts := strings.Split(lineParts[1], "|")
		winningNums := make([]int, 0)
		for _, numStr := range strings.Fields(numsParts[0]) {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to parse %s: %w", numStr, err)
			}
			winningNums = append(winningNums, num)
		}
		existingNums := make([]int, 0)
		for _, numStr := range strings.Fields(numsParts[1]) {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to parse %s: %w", numStr, err)
			}
			existingNums = append(existingNums, num)
		}
		winningCards = append(winningCards, winningNums)
		existingCards = append(existingCards, existingNums)
	}
	return winningCards, existingCards, nil
}
