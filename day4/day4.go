package main

import (
	"fmt"
	"slices"

	"aoc/util"
)

func convertLinesTo2d(lines []string) ([][]rune, error) {
	result := [][]rune{}
	for i := range lines {
		result = append(result, []rune(lines[i]))
	}

	if len(result) == 0 {
		return result, fmt.Errorf("No lines found")
	}

	return result, nil
}

func debugPointInGrid(grid [][]rune, startRow, startCol int) {
	topLeft, _ := util.GridAt(grid, startRow-1, startCol-1)
	top, _ := util.GridAt(grid, startRow-1, startCol)
	topRight, _ := util.GridAt(grid, startRow-1, startCol+1)
	right, _ := util.GridAt(grid, startRow, startCol+1)
	center, _ := util.GridAt(grid, startRow, startCol)
	bottomRight, _ := util.GridAt(grid, startRow+1, startCol+1)
	bottom, _ := util.GridAt(grid, startRow+1, startCol)
	bottomLeft, _ := util.GridAt(grid, startRow+1, startCol-1)
	left, _ := util.GridAt(grid, startRow, startCol-1)

	fmt.Printf("%q%q%q\n", topLeft, top, topRight)
	fmt.Printf("%q%q%q\n", left, center, right)
	fmt.Printf("%q%q%q\n", bottomLeft, bottom, bottomRight)
}

type word struct {
	target    string
	column    int
	row       int
	direction string
}

var directionMap = map[string][]int{
	"UpLeft":      {-1, -1},
	"Up":          {-1, 0},
	"UpRight":     {-1, 1},
	"Left":        {0, -1},
	"Right":       {0, 1},
	"BottomLeft":  {1, -1},
	"Bototm":      {1, 0},
	"BottomRight": {1, 1},
}

func findWord(grid [][]rune, target string) []word {
	targetRunes := []rune(target)
	words := make([]word, 0)

	for row := range len(grid) {
		for col := range len(grid[0]) {
			// Always start by looking for the first letter in the target string
			if grid[row][col] == targetRunes[0] {
				// Now look around in each of the directions.
				for directionName, dir := range directionMap {
					// For each direction, search in that direction, using the target runes
					// if we find a match, add it to the words slice.
					if match := searchDirection(grid, row, col, dir, targetRunes); match {
						words = append(words, word{
							target:    target,
							row:       row,
							column:    col,
							direction: directionName,
						})
					}
				}
			}
		}
	}
	return words
}

func searchDirection(grid [][]rune, startRow, startCol int, dir []int, target []rune) bool {
	targetLength := len(target)

	// Direction is either 1, 0 or -1 in a direction, so times that by number of characters
	// in the target string to get the possible end of it.
	endRow := startRow + dir[0]*(targetLength-1)
	endCol := startCol + dir[1]*(targetLength-1)

	// Make sure the end row and col are within bounds of the grid
	if endRow < 0 || endRow >= len(grid) || endCol < 0 || endCol >= len(grid[0]) {
		return false
	}

	// We know it can fit inside the grid now, check each target
	// rune one by one in the direction provided.
	for i := range target {
		currentRow := startRow + dir[0]*i
		currentCol := startCol + dir[1]*i

		// As soon as we don't find a match, stop searching.
		if grid[currentRow][currentCol] != target[i] {
			return false
		}
	}
	return true
}

func findXMas(grid [][]rune) int {
	finds := 0
	targetRune := 'A'

	for row := range len(grid) {
		for col := range len(grid[0]) {
			if grid[row][col] == targetRune {
				result := searchForX(grid, row, col)
				if result {
					finds++
				}
			}
		}
	}
	return finds
}

func searchForX(grid [][]rune, startRow, startCol int) bool {
	// Check each of the corners of the start position
	topLeft, topLeftErr := util.GridAt(grid, startRow-1, startCol-1)
	if topLeftErr != nil {
		return false
	}
	topRight, topRightErr := util.GridAt(grid, startRow-1, startCol+1)
	if topRightErr != nil {
		return false
	}
	bottomRight, bottomRightErr := util.GridAt(grid, startRow+1, startCol+1)
	if bottomRightErr != nil {
		return false
	}
	bottomLeft, bottomLeftErr := util.GridAt(grid, startRow+1, startCol-1)
	if bottomLeftErr != nil {
		return false
	}

	// Check if the assembled string in both diagonal directions
	// match either of the valid values
	topLeft_bottomRight := fmt.Sprintf("%c%c%c", topLeft, 'A', bottomRight)
	topRight_bottomLeft := fmt.Sprintf("%c%c%c", topRight, 'A', bottomLeft)
	valid := []string{"MAS", "SAM"}

	return slices.Contains(valid, topLeft_bottomRight) && slices.Contains(valid, topRight_bottomLeft)
}

func main() {
	success, lines := util.ReadInputIntoLines("input.txt")
	if success {
		target := "XMAS"
		arr, _ := convertLinesTo2d(lines)
		words := findWord(arr, target)

		fmt.Printf("Part1: Found %d number of words in grid\n", len(words))

		count := findXMas(arr)
		fmt.Printf("Part2: Found %d number of X-shaped MAS in grid", count)
	}
}
