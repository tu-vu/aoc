package day02

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var content string // Embed file content into string

func cubeConundrumPart1(input string) int {
	var sum int
	maxCubesPerGame := make(map[int]map[string]float64)
	err := parseInput(input, maxCubesPerGame)
	if err != nil {
		panic(err)
	}
	for gameID, maxCubes := range maxCubesPerGame {
		if maxCubes["red"] <= 12 && maxCubes["green"] <= 13 && maxCubes["blue"] <= 14 {
			sum += gameID
		}
	}
	return sum
}

func cubeConundrumPart2(input string) float64 {
	var sum float64
	maxCubesPerGame := make(map[int]map[string]float64)
	err := parseInput(input, maxCubesPerGame)
	if err != nil {
		panic(err)
	}
	for _, maxCubes := range maxCubesPerGame {
		sum += maxCubes["red"] * maxCubes["green"] * maxCubes["blue"]
	}
	return sum
}

func parseInput(input string, maxCubesPerGame map[int]map[string]float64) error {
	for _, line := range strings.Split(input, "\n") {
		lineParts := strings.Split(line, ":")
		gameID, err := strconv.Atoi(lineParts[0][5:])
		if err != nil {
			return fmt.Errorf("invalid game ID: %s", lineParts[0][5:])
		}
		maxCubesPerGame[gameID] = make(map[string]float64)
		sets := strings.Split(strings.Trim(lineParts[1], " "), "; ")
		for _, set := range sets {
			setParts := strings.Split(set, ", ")
			for _, part := range setParts {
				countAndColor := strings.Split(part, " ")
				count, err := strconv.ParseFloat(countAndColor[0], 64)
				if err != nil {
					return fmt.Errorf("invalid count: %s", countAndColor[0])
				}
				color := countAndColor[1]
				maxCubesPerGame[gameID][color] = math.Max(maxCubesPerGame[gameID][color], count)
			}
		}
	}
	return nil
}
