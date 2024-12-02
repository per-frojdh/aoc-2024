package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"aoc/util"
)

func getDistance(a, b int) int {
	result := a - b
	if result < 0 {
		return result * -1
	}

	return result
}

func splitNumbers(lines []string) ([]int, []int) {
	left := []int{}
	right := []int{}

	for i := range len(lines) {
		current := lines[i]
		s := strings.Fields(current)
		l, err := strconv.Atoi(s[0])
		if err != nil {
			fmt.Println("s[0]", s[0])
		}
		left = append(left, l)

		r, err := strconv.Atoi(s[1])
		if err != nil {
			fmt.Println("s[1]", s[1])
		}
		right = append(right, r)
	}
	return left, right
}

func sortNumbers(left, right []int) ([]int, []int) {
	sort.Ints(left)
	sort.Ints(right)
	return left, right
}

func findOccurrences(numbers []int, target int) int {
	occurrences := 0

	for i := range numbers {
		if numbers[i] == target {
			occurrences += 1
		}
	}

	return occurrences
}

func calculateDistance(lines []string) int {
	distance := 0
	left, right := splitNumbers(lines)
	left, right = sortNumbers(left, right)

	for i := range len(left) {
		distance += getDistance(left[i], right[i])
	}

	return distance
}

func getSimilarityScore(lines []string) int {
	similarityScore := 0
	left, right := splitNumbers(lines)

	for i := range len(left) {
		current := left[i]
		occurences := findOccurrences(right, current)
		similarityScore += current * occurences
	}

	return similarityScore
}

func main() {
	success, lines := util.ReadInputIntoLines("input.txt")
	if success {
		distance := calculateDistance(lines)
		fmt.Printf("Part1: Distance is %d \n", distance)

		score := getSimilarityScore(lines)
		fmt.Printf("Part2: Score is %d \n", score)
	}
}
