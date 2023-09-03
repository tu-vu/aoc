package d1_calorie_counting

import (
	"fmt"
	"testing"
)

type testcase struct {
	inputFile string
}

func TestCalorieCountingPart1(t *testing.T) {
	testcases := []testcase{
		{
			inputFile: "input1.txt",
		},
	}
	fmt.Printf("--------- Day 1: Calorie Counting, Part 1 ----------\n")
	for _, tc := range testcases {
		calo := calorieCountingPart1(tc.inputFile)
		fmt.Printf("[inputFile] :%v, [gotCalo]: %v", tc.inputFile, calo)
	}
	fmt.Printf("\n-----------------------------------------------------\n")
}

func TestCalorieCountingPart2(t *testing.T) {
	testcases := []testcase{
		{
			inputFile: "input1.txt",
		},
	}
	fmt.Printf("--------- Day 1: Calorie Counting, Part 2 ----------\n")
	for _, tc := range testcases {
		calo := calorieCountingPart2(tc.inputFile)
		fmt.Printf("[inputFile] :%v, [gotCalo]: %v", tc.inputFile, calo)
	}
	fmt.Printf("\n-----------------------------------------------------\n")
}
