package day02

import (
	_ "embed"
	"fmt"
	"strings"
)

// roundScores[i][j] map score for a round when player 1 plays i and player 2 plays j
var roundScores = map[string]map[string]int{
	"A": {"X": 3, "Y": 6, "Z": 0},
	"B": {"X": 0, "Y": 3, "Z": 6},
	"C": {"X": 6, "Y": 0, "Z": 3},
}

// shapeScores[i] map score for a round when shape i is played
// For part 1, both 'A' and 'X' et al. are considered shape.
// For part 2, only 'A', 'B', 'C' are considered shape.
var shapeScores = map[string]int{
	"A": 1, "X": 1,
	"B": 2, "Y": 2,
	"C": 3, "Z": 3,
}

// ruleScores[i] map score for a round given the following rules:
// X -> Lose (0)
// Y -> Draw (3)
// Z -> Win  (6)
var ruleScores = map[string]int{"X": 0, "Y": 3, "Z": 6}

// targetShapes[i][j] map the shape needed from player 2 to follow the rules
var targetShapes = map[string]map[string]string{
	"A": {"X": "C", "Y": "A", "Z": "B"},
	"B": {"X": "A", "Y": "B", "Z": "C"},
	"C": {"X": "B", "Y": "C", "Z": "A"},
}

//go:embed input.txt
var content string

func rockPaperScissorsPart1(input string) (int, error) {
	var totalScore int
	for _, line := range strings.Split(input, "\n") {
		// Shape each player selected e.g. ["A", "X"]
		player := strings.Split(line, " ")
		if len(player) != 2 {
			return totalScore, fmt.Errorf("Invalid input: %v\n", line)
		}

		// Calculate score for player 2
		roundScore, ok := roundScores[player[0]][player[1]]
		if !ok {
			return totalScore, fmt.Errorf("Round score: Invalid shape, player 1 %v, player 2 %v\n", player[0], player[1])
		}
		totalScore += roundScore

		shapeScore, ok := shapeScores[player[1]]
		if !ok {
			return totalScore, fmt.Errorf("Shape score: Invalid shape, player 2 %v\n", player[1])
		}
		totalScore += shapeScore
	}
	return totalScore, nil
}

func rockPaperScissorsPart2(input string) (int, error) {
	var totalScore int
	for _, line := range strings.Split(input, "\n") {
		// Shape each player selected e.g. ["A", "X"]
		player := strings.Split(line, " ")
		if len(player) != 2 {
			return totalScore, fmt.Errorf("Invalid input: %v\n", line)
		}

		// Calculate score for player 2
		ruleScore, ok := ruleScores[player[1]]
		if !ok {
			return totalScore, fmt.Errorf("Rule score: Invalid shape, player 2 %v\n", player[1])
		}
		totalScore += ruleScore

		// Get the shape needed from player 2 to follow the rules
		targetShape, ok := targetShapes[player[0]][player[1]]
		if !ok {
			return totalScore, fmt.Errorf("Target shape: Invalid shape, player 1 %v player 2 %v\n", player[0], player[1])
		}

		shapeScore, ok := shapeScores[targetShape]
		if !ok {
			return totalScore, fmt.Errorf("Shape score: Invalid shape, player 2 %v\n", player[1])
		}
		totalScore += shapeScore
	}
	return totalScore, nil
}
