package main

import (
	"fmt"
	"testing"

	"aoc/util"
)

func createTestinput() []string {
	_, lines := util.ReadInputIntoLines("test_input.txt")
	return lines
}

func createPartTwoTestInput() []string {
	_, lines := util.ReadInputIntoLines("test_input_pt2.txt")
	return lines
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

func debugWords(words []word) {
	for i := range words {
		word := words[i]
		fmt.Printf("Row: %d, Col: %d, Direction: %q\n", word.row+1, word.column+1, word.direction)
	}
}

func debugRunes(runes [][]rune) {
	for i := range runes {
		line := runes[i]
		fmt.Println(string(line))
	}
}

func TestConvertLinesTo2d(t *testing.T) {
	t.Run("Can convert test input into 2d", func(t *testing.T) {
		lines := createTestinput()
		arr, err := convertLinesTo2d(lines)

		if err != nil {
			t.Errorf("got an error from convertLinesTo2d: %q", err)
		}

		expectedLength := 10
		expectedDepth := 10

		if len(arr) != expectedLength {
			t.Errorf("got %d, want %d", len(arr), expectedLength)
		}

		for i := range arr {
			if len(arr[i]) != expectedDepth {
				t.Errorf("got %d, want %d at index %d", len(arr[i]), expectedDepth, i)
			}
		}
	})
}

func TestSearchForMatchingStringsIn2d(t *testing.T) {
	t.Run("Can find XMAS inside testing string", func(t *testing.T) {
		lines := createTestinput()
		arr, _ := convertLinesTo2d(lines)
		expectedLength := 18
		words := findWord(arr, "XMAS")

		if len(words) != expectedLength {
			t.Errorf("got %d, want %d", len(words), expectedLength)
		}
	})
}

func TestSearchForXMasIn2d(t *testing.T) {
	t.Run("Can find an X-shaped MAS inside testing string", func(t *testing.T) {
		lines := createPartTwoTestInput()
		arr, _ := convertLinesTo2d(lines)

		expectedLength := 9
		count := findXMas(arr)

		if count != expectedLength {
			t.Errorf("got %d, want %d", count, expectedLength)
		}
	})
}
