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
