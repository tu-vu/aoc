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
	cardToMatchesCount := make(map[int]int)
	for i := 0; i < len(winningCards); i++ {
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
		// Map the card to its number of matches
		cardToMatchesCount[i] = count
	}
	// dp[i] is the number of copies required to process the ith card
	dp := make([]int, len(winningCards))
	for i := len(winningCards) - 1; i >= 0; i-- {
		// Every card starts with 1 copy of itself
		dp[i] = 1
		// Update cached value with the number of copies of the cards that it matches
		for j := i + 1; j < i+1+cardToMatchesCount[i]; j++ {
			dp[i] += dp[j]
		}
	}
	var sum int
	for i := 0; i < len(dp); i++ {
		sum += dp[i]
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
