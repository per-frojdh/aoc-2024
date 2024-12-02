package main

import (
	"fmt"
	"testing"
)

func createRedditinput() []string {
	_, lines := readInputIntoLines("reddit_input.txt")
	return lines
}

func createTestinput() []string {
	_, lines := readInputIntoLines("test_input.txt")
	return lines
}

func assertSlices(t *testing.T, a, b []int) {
	t.Helper()

	for i := range len(a) {
		if a[i] != b[i] {
			t.Errorf("Expected %d and %d to be equal", a[i], b[i])
		}
	}
}

func TestReportParsing(t *testing.T) {
	t.Run("Can parse a single line of a report to levels", func(t *testing.T) {
		lines := createTestinput()
		first := lines[0]

		levels := parseReport(first)
		want := []int{7, 6, 4, 2, 1}

		assertSlices(t, want, levels)
	})

	t.Run("Can parse a single line of a report to levels", func(t *testing.T) {
		lines := createTestinput()
		reports := parseReports(lines)

		first := reports[0]
		last := reports[len(reports)-1]

		assertSlices(t, []int{7, 6, 4, 2, 1}, first)
		assertSlices(t, []int{1, 3, 6, 7, 9}, last)
	})
}

func TestLevelLogic(t *testing.T) {
	lines := createTestinput()
	var testExampleMap = map[int]bool{
		0: true,
		1: false,
		2: false,
		3: false,
		4: false,
		5: true,
	}

	for i := range len(lines) {
		t.Run(fmt.Sprintf("Checking if example level %d passes safety", i+1), func(t *testing.T) {
			current := lines[i]
			levels := parseReport(current)
			want := testExampleMap[i]

			fmt.Println(levels)
			got := GetSafety(levels)

			if got != want {
				t.Errorf("got %t want %t", got, want)
			}
		})
	}
}

func TestLevelLogicPartTwo(t *testing.T) {
	lines := createTestinput()
	var testExampleMap = map[int]bool{
		0: true,
		1: false,
		2: false,
		3: true,
		4: true,
		5: true,
	}

	for i := range len(lines) {
		t.Run(fmt.Sprintf("Checking if example level %d passes safety (part 2)", i+1), func(t *testing.T) {
			current := lines[i]
			levels := parseReport(current)
			want := testExampleMap[i]

			got := GetSafetyWithProblemDampener(levels)

			if got != want {
				t.Errorf("got %t want %t", got, want)
			}
		})
	}
}

func TestReddit(t *testing.T) {
	lines := createRedditinput()
	var testExampleMap = map[int]bool{
		0:  true,
		1:  true,
		2:  true,
		3:  true,
		4:  true,
		5:  true,
		6:  true,
		7:  true,
		8:  true,
		9:  true,
		10: true,
	}

	for i := range len(lines) {
		t.Run(fmt.Sprintf("Checking if example level %d passes safety (part 2)", i+1), func(t *testing.T) {
			current := lines[i]
			levels := parseReport(current)
			want := testExampleMap[i]

			got := GetSafetyWithProblemDampener(levels)

			if got != want {
				t.Errorf("got %t want %t", got, want)
			}
		})
	}
}
