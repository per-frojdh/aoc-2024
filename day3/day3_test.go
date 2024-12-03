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

func createTestPartTwoInput() string {
	_, lines := util.ReadInputIntoString("test_input_pt2.txt")
	return lines
}

func TestParsingCorruptedInput(t *testing.T) {
	t.Run("Can parse the first example input", func(t *testing.T) {
		input := createTestinput()
		firstLine := input[0]
		expectedMatches := 4

		_, parsed := parseCorruptedInput(firstLine)

		if len(parsed) != expectedMatches {
			t.Errorf("got %d, want %d", len(parsed), expectedMatches)
		}
	})

	t.Run("Can parse all example inputs", func(t *testing.T) {
		input := createTestinput()
		allMatches := make([][]int, 0)
		expectedMatches := 5

		for i := range input {
			success, parsed := parseCorruptedInput(input[i])
			if success {
				allMatches = append(allMatches, parsed...)
			}
		}

		if len(allMatches) != expectedMatches {
			fmt.Println(allMatches)
			t.Errorf("got %d, want %d", len(allMatches), expectedMatches)
		}
	})
}

func TestComputingCorruptedInputs(t *testing.T) {
	t.Run("Can compue the result of first parsed set of inputs", func(t *testing.T) {
		input := createTestinput()
		firstLine := input[0]
		expected := 161

		_, parsed := parseCorruptedInput(firstLine)

		result := computeCorruptedInput(parsed)

		if result != expected {
			t.Errorf("got %d, want %d", result, expected)
		}
	})
}

func TestPartTwo(t *testing.T) {
	t.Run("Can parse the first example input", func(t *testing.T) {
		input := createTestPartTwoInput()
		expectedMatches := 2

		parsed := parseCorruptedWithExtraInstructions(input)

		if len(parsed) != expectedMatches {
			t.Errorf("got %d, want %d", len(parsed), expectedMatches)
		}
	})
}
