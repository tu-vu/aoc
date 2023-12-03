package day03

import (
	"aoc/util"
	_ "embed"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var content string // Embed file content into string

func gearRatiosPart1(input string) int {
	sum := 0
	mat := parseInput(input)
	for i := 0; i < len(mat); i++ {
		numStr := ""
		isAdjacentToSymbol := false
		for j := 0; j <= len(mat[i]); j++ {
			if j < len(mat[i]) && unicode.IsDigit(mat[i][j]) {
				numStr += string(mat[i][j])
				isAdjacentToSymbol = isAdjacentToSymbol || isAdjecentToSymbol(mat, i, j)
			} else if numStr != "" {
				if isAdjacentToSymbol {
					val, err := strconv.Atoi(numStr)
					if err != nil {
						panic(err)
					}
					sum += val
				}
				numStr = ""
				isAdjacentToSymbol = false
			}
		}
	}
	return sum
}

func parseInput(input string) [][]rune {
	mat := make([][]rune, 0)
	for _, line := range strings.Split(input, "\n") {
		row := make([]rune, 0)
		for _, r := range line {
			row = append(row, r)
		}
		mat = append(mat, row)
	}
	return mat
}

func isAdjecentToSymbol(mat [][]rune, i int, j int) bool {
	dirs := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1}, {1, -1},
		{1, 0}, {1, 1}, {0, -1}, {0, 1},
	}
	for _, dir := range dirs {
		r, c := i+dir[0], j+dir[1]
		if 0 <= r && r < len(mat) && 0 <= c && c < len(mat[i]) &&
			!unicode.IsDigit(mat[r][c]) && mat[r][c] != '.' {
			return true
		}
	}
	return false
}

// Returns an array of pair i, j which are the locations of the adjacent numbers
func getAdjacentNumbers(mat [][]rune, numberLocs map[int]map[int]int, i int, j int) map[int]struct{} {
	// TODO: Add set implementation
	numberSet := make(map[int]struct{})
	dirs := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1}, {1, -1},
		{1, 0}, {1, 1}, {0, -1}, {0, 1},
	}
	for _, dir := range dirs {
		r, c := i+dir[0], j+dir[1]
		if 0 <= r && r < len(mat) && 0 <= c && c < len(mat[i]) && unicode.IsDigit(mat[r][c]) {
			if val, ok := numberLocs[r][c]; ok {
				_, ok := numberSet[val]
				if !ok {
					numberSet[val] = struct{}{}
				}
			}
		}
	}
	return numberSet
}

func gearRatiosPart2(input string) int {
	sum := 0
	mat := parseInput(input)

	// 1. Find the locations of the numbers
	numberLocs := make(map[int]map[int]int)
	for i := 0; i < len(mat); i++ {
		numStr := ""
		startIndex := -1
		for j := 0; j <= len(mat[i]); j++ {
			if j < len(mat[i]) && unicode.IsDigit(mat[i][j]) {
				numStr += string(mat[i][j])
				if startIndex == -1 {
					startIndex = j
				}
			} else if numStr != "" {
				// Get the number
				val, err := strconv.Atoi(numStr)
				if err != nil {
					panic(err)
				}
				// Initialize map if not exist
				if _, ok := numberLocs[i]; !ok {
					numberLocs[i] = make(map[int]int)
				}
				// Map the location of number to its value
				for k := startIndex; k < util.Min(j, len(mat[i])); k++ {
					numberLocs[i][k] = val
				}
				numStr = ""
				startIndex = -1
			}
		}
	}

	// 2. Find the locations of the gears
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == '*' {
				adjNumbers := getAdjacentNumbers(mat, numberLocs, i, j)
				if len(adjNumbers) == 2 {
					mul := 1
					for num := range adjNumbers {
						mul *= num
					}
					sum += mul
				}
			}
		}
	}
	return sum
}
