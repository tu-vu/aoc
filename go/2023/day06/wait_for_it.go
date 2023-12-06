package day06

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var content string // Embed file content into string

func waitForItPart1(input string) float64 {
	var ans float64 = 1
	time, distance, err := parseInputPart1(input)
	if err != nil {
		panic(err)
	}
	// For each race, find the min/max h such that h * (t - h) >= d given 0 <= h <= t
	// No of ways to beat a race = maxH - minH + 1
	for i, t := range time {
		minH, maxH := findMinWaysToBeat(0, t, t, distance[i]), findMaxWaysToBeat(0, t, t, distance[i])
		ans *= maxH - minH + 1
	}
	return ans
}

func parseInputPart1(input string) ([]int, []int, error) {
	lines := strings.Split(input, "\n")
	time := make([]int, 0)
	timeParts := strings.Fields(lines[0])[1:]
	for i := 0; i < len(timeParts); i++ {
		timeInt, err := strconv.Atoi(timeParts[i])
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse %s: %w", timeParts[i], err)
		}
		time = append(time, timeInt)
	}
	distance := make([]int, 0)
	distanceParts := strings.Fields(lines[1])[1:]
	for i := 0; i < len(distanceParts); i++ {
		distanceInt, err := strconv.Atoi(distanceParts[i])
		if err != nil {
			return nil, nil, fmt.Errorf("failed to parse %s: %w", distanceParts[i], err)
		}
		distance = append(distance, distanceInt)
	}
	return time, distance, nil
}

func waitForItPart2(input string) float64 {
	time, distance, err := parseInputPart2(input)
	if err != nil {
		panic(err)
	}
	minH, maxH := findMinWaysToBeat(0, time, time, distance), findMaxWaysToBeat(0, time, time, distance)
	return maxH - minH + 1
}

func parseInputPart2(input string) (int, int, error) {
	lines := strings.Split(input, "\n")
	timeParts := strings.Fields(lines[0])[1:]
	timeStr := strings.Join(timeParts, "")
	time, err := strconv.Atoi(timeStr)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse %s: %w", timeStr, err)
	}
	distanceParts := strings.Fields(lines[1])[1:]
	distanceStr := strings.Join(distanceParts, "")
	distance, err := strconv.Atoi(distanceStr)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse %s: %w", distanceStr, err)
	}
	return time, distance, nil
}

func findMinWaysToBeat(lo int, hi int, time int, distance int) float64 {
	if lo > hi {
		return math.Inf(1)
	}
	mid := lo + (hi-lo)/2
	if mid*(time-mid) > distance {
		// mid is a valid hold time, but check if there's smaller hold time in the lower half
		return math.Min(float64(mid), findMinWaysToBeat(lo, mid-1, time, distance))
	} else {
		// Min hold time is in the upper half
		return findMinWaysToBeat(mid+1, hi, time, distance)
	}
}

func findMaxWaysToBeat(lo int, hi int, time int, distance int) float64 {
	if lo > hi {
		return math.Inf(-1)
	}
	mid := lo + (hi-lo)/2
	if mid*(time-mid) > distance {
		// mid is a valid hold time, but check if there's larger hold time in the upper half
		return math.Max(float64(mid), findMaxWaysToBeat(mid+1, hi, time, distance))
	} else {
		// Max hold time is in the lower half
		return findMaxWaysToBeat(lo, mid-1, time, distance)
	}
}
