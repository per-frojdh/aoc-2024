package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"aoc/util"
)

const (
	Increasing = 1
	NotMoving  = 0
	Decreasing = -1
)

func parseReport(report string) []int {
	fields := strings.Fields(report)
	levels := []int{}
	for i := range len(fields) {
		current, _ := strconv.Atoi(fields[i])
		levels = append(levels, current)
	}

	return levels
}

func parseReports(report []string) [][]int {
	reports := make([][]int, len(report))

	for i := range len(report) {
		reports[i] = parseReport(report[i])
	}

	return reports
}

func getDistance(a, b int) int {
	result := a - b
	if result < 0 {
		return result * -1
	}

	return result
}

func getDirection(x int) int {
	switch {
	case x < 0:
		return Increasing
	case x > 0:
		return Decreasing
	default:
		return NotMoving
	}
}

func GetSafetyWithProblemDampener(report []int) bool {
	if GetSafety(report) {
		return true
	}

	for i := range report {
		// Brute-force check each value in the level by removing it
		// and checking if it works this time, until it does (or doesnt)
		portion := slices.Delete(slices.Clone(report), i, i+1)
		if GetSafety(portion) {
			return true
		}
	}
	return false
}

func GetSafety(levels []int) bool {
	// Get initial direction by looking at first two values
	direction := getDirection(levels[0] - levels[1])
	for i := range len(levels) - 1 {
		// If we change direction, we are unsafe
		changedDirection := getDirection(levels[i] - levels[i+1])
		if changedDirection != direction {
			return false
		}

		// If the distance is less than 1 or greater than 3, we are unsafe
		distance := getDistance(levels[i], levels[i+1])
		if distance < 1 || distance > 3 {
			return false
		}
	}

	return true

}

func main() {
	success, lines := util.ReadInputIntoLines("input.txt")
	if success {
		reports := parseReports(lines)
		numberOfSafeReports_pt1 := 0
		numberOfSafeReports_pt2 := 0

		for i := range len(reports) {
			safe := GetSafety(reports[i])
			if safe {
				numberOfSafeReports_pt1 += 1
			}
		}

		for i := range len(reports) {
			safe := GetSafetyWithProblemDampener(reports[i])
			if safe {
				numberOfSafeReports_pt2 += 1
			}
		}
		fmt.Printf("Part1: Number of safe states are: %d \n", numberOfSafeReports_pt1)
		fmt.Printf("Part2: Number of safe states are: %d \n", numberOfSafeReports_pt2)
	}
}
