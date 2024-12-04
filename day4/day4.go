package main

import (
	"fmt"

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

type word struct {
	target    string
	column    int
	row       int
	direction string
}

var directionMap = map[string][]int{
	"UpLeft":    {-1, -1},
	"Up":        {-1, 0},
	"UpRight":   {-1, 1},
	"Left":      {0, -1},
	"Right":     {0, 1},
	"DownLeft":  {1, -1},
	"Down":      {1, 0},
	"DownRight": {1, 1},
}

func findWord(grid [][]rune, target string) []word {
	if len(grid) == 0 {
		return nil
	}
	targetRunes := []rune(target)
	words := make([]word, 0)

	for row := range len(grid) {
		for col := range len(grid[0]) {
			// If we find the first letter, check all directions
			if grid[row][col] == targetRunes[0] {
				for directionName, dir := range directionMap {
					if match := searchDirection(grid, row, col, dir, targetRunes); match != nil {
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

func findX(grid [][])

func searchForX(grid [][]rune, startRow, startCol int) [][2]int {
	rows := len(grid)
	cols := len(grid[0])
	target := "MAS"
	path := make([][2]int, len(target))

	// Check if the word would fit in this direction
	endRow := startRow + len(target)-1
	endCol := startCol + len(target)-1

	if endRow < 0 || endRow >= rows || endCol < 0 || endCol >= cols {
		return nil
	}
	
	// M.S
	// .A.
	// M.S
	// We have found a A
	
	// startRow -1, startCol -1 HAS to be M
	// startRow +1  startCol -1 has to be M
	// startRow -1 startCol +1 has to be S
	// startRow +1 startCol +1 has to be S


	return path
}

func searchDirection(grid [][]rune, startRow, startCol int, dir []int, target []rune) [][2]int {
	rows := len(grid)
	cols := len(grid[0])
	path := make([][2]int, len(target))

	// Check if the word would fit in this direction
	endRow := startRow + dir[0]*(len(target)-1)
	endCol := startCol + dir[1]*(len(target)-1)

	if endRow < 0 || endRow >= rows || endCol < 0 || endCol >= cols {
		return nil
	}

	// Check each character in the direction
	for i := 0; i < len(target); i++ {
		currentRow := startRow + dir[0]*i
		currentCol := startCol + dir[1]*i

		if grid[currentRow][currentCol] != target[i] {
			return nil
		}
		path[i] = [2]int{currentRow, currentCol}
	}

	return path
}

func main() {
	success, lines := util.ReadInputIntoLines("input.txt")
	if success {
		target := "XMAS"
		arr, _ := convertLinesTo2d(lines)
		words := findWord(arr, target)

		fmt.Printf("Part1: Found %d number of words in grid", len(words))
	}
}
