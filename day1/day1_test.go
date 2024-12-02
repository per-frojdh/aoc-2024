package main

import (
	"testing"

	"aoc/util"
)

func createTestinput() []string {
	_, lines := util.ReadInputIntoLines("test_input.txt")
	return lines
}

func TestReadInputIntoLines(t *testing.T) {
	t.Run("can read test-input", func(t *testing.T) {
		lines := createTestinput()
		firstLine := lines[len(lines)-1]

		want := "3   3\r\n"

		if firstLine != want {
			t.Errorf("got %q want %q", firstLine, want)
		}
	})
}

func TestSplitAndSortNumbers(t *testing.T) {
	lines := createTestinput()

	t.Run("will find lowest numbers", func(t *testing.T) {
		left, right := splitNumbers(lines)
		left, right = sortNumbers(left, right)

		firstLeft := left[0]
		firstRight := right[0]

		wantLeft := 1
		wantRight := 3

		if firstLeft != wantLeft {
			t.Errorf("got %q want %q", firstLeft, wantLeft)
		}

		if firstRight != wantRight {
			t.Errorf("got %q want %q", firstRight, wantRight)
		}
	})

	t.Run("will find highest numbers", func(t *testing.T) {
		left, right := splitNumbers(lines)
		left, right = sortNumbers(left, right)

		lastLeft := left[len(left)-1]
		lastRight := right[len(right)-1]

		wantLeft := 4
		wantRight := 9

		if lastLeft != wantLeft {
			t.Errorf("got %q want %q", lastLeft, wantLeft)
		}

		if lastRight != wantRight {
			t.Errorf("got %q want %q", lastRight, wantRight)
		}
	})
}

func TestDistance(t *testing.T) {
	lines := createTestinput()
	left, right := splitNumbers(lines)
	left, right = sortNumbers(left, right)

	t.Run("Distance between lowest set of numbers", func(t *testing.T) {
		firstLeft := left[0]
		firstRight := right[0]

		wantDistance := 2

		distance := getDistance(firstLeft, firstRight)

		if distance != wantDistance {
			t.Errorf("got %q want %q", distance, wantDistance)
		}
	})

	t.Run("Distance between equal set of numbers", func(t *testing.T) {
		firstLeft := left[2]
		firstRight := right[2]

		wantDistance := 0

		distance := getDistance(firstLeft, firstRight)

		if distance != wantDistance {
			t.Errorf("got %q want %q", distance, wantDistance)
		}
	})

	t.Run("Distance that would produce a negative value", func(t *testing.T) {
		distance := getDistance(3, 2)
		previousDistance := 5
		wantDistance := 1

		if distance != wantDistance {
			t.Errorf("got %d want %d", distance, wantDistance)
		}

		previousDistance += distance
		expectedDistance := 6

		if previousDistance != expectedDistance {
			t.Errorf("got %d want %d", previousDistance, expectedDistance)
		}
	})
}

func TestCalculateDistance(t *testing.T) {
	t.Run("Can calculate distance of test_input", func(t *testing.T) {
		lines := createTestinput()
		distance := calculateDistance(lines)
		want := 11

		if distance != want {
			t.Errorf("got %d want %d", distance, want)
		}
	})
}

func TestGetSimilarityScore(t *testing.T) {

	t.Run("Can find occurrences in number array", func(t *testing.T) {
		lines := createTestinput()
		_, right := splitNumbers(lines)

		occurrences := findOccurrences(right, 3)
		want := 3

		if occurrences != want {
			t.Errorf("got %d want %d", occurrences, want)
		}
	})
	t.Run("Can get similarity score from test input", func(t *testing.T) {
		lines := createTestinput()
		score := getSimilarityScore(lines)

		want := 31

		if score != want {
			t.Errorf("got %d want %d", score, want)
		}
	})
}
