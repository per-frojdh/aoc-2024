package main

import (
	"aoc/util"
	"fmt"
	"slices"
	"testing"
)

func createTestinput() []string {
	_, lines := util.ReadInputIntoLines("test_input.txt")
	return lines
}

func createTestInputPartTwo() []string {
	_, lines := util.ReadInputIntoLines("test_input_pt2.txt")
	return lines
}

func TestSeparationOfRulesAndUpdates(t *testing.T) {
	t.Run("Can successfully extract rules and updates", func(t *testing.T) {
		lines := createTestinput()
		expectedRulesLength := 21
		expectedUpdatesLength := 6
		rules, updates := separateRulesAndUpdates(lines)

		if len(rules) != expectedRulesLength {
			t.Errorf("got %d, want %d", len(rules), expectedRulesLength)
		}

		if len(updates) != expectedUpdatesLength {
			t.Errorf("got %d, want %d", len(updates), expectedUpdatesLength)
		}

		for i := range updates {
			if len(updates[i]) < 1 {
				t.Errorf("got %d, want %d", len(updates[i]), 0)
			}
		}
	})
}

func TestCheckCorrectness(t *testing.T) {
	var exampleResults = map[int]bool{
		0: true,
		1: true,
		2: true,
		3: false,
		4: false,
		5: false,
	}

	for idx, expected := range exampleResults {
		lines := createTestinput()
		t.Run(fmt.Sprintf("Can check rules for update %d", idx), func(t *testing.T) {
			rules, updates := separateRulesAndUpdates(lines)
			updateToCheck := updates[idx]
			got := checkCorrectness(updateToCheck, rules)

			if got != expected {
				// fmt.Println(updates[idx])
				t.Errorf("got %t, want %t on update with index %d", got, expected, idx)
			}

		})
	}
}

func TestCountingMiddles(t *testing.T) {

	t.Run("Can get total amount of all correct middles", func(t *testing.T) {
		lines := createTestinput()
		rules, updates := separateRulesAndUpdates(lines)
		result := checkUpdatesAgainstRules(rules, updates)

		total := 0
		expected := 143
		for _, count := range result {
			total += count
		}

		if total != expected {
			t.Errorf("got %d, want %d", total, expected)
		}

	})
}

func TestSortingForPartTwo(t *testing.T) {

	t.Run("Can sort first example output", func(t *testing.T) {
		lines := createTestInputPartTwo()
		rules, updates := separateRulesAndUpdates(lines)
		first := updates[0]
		sorted := fixSorting(first, rules)
		expected := []int{97, 75, 47, 61, 53}

		if !slices.Equal(sorted, expected) {
			t.Errorf("not equal")
			fmt.Println("sorted", sorted)
			fmt.Println("expected", expected)
		}
	})

	t.Run("Can sort second example output", func(t *testing.T) {
		lines := createTestInputPartTwo()
		rules, updates := separateRulesAndUpdates(lines)
		second := updates[1]
		sorted := fixSorting(second, rules)
		expected := []int{61, 29, 13}

		if !slices.Equal(sorted, expected) {
			t.Errorf("not equal")
			fmt.Println("sorted", sorted)
			fmt.Println("expected", expected)
		}
	})

	t.Run("Can get third example output", func(t *testing.T) {
		lines := createTestInputPartTwo()
		rules, updates := separateRulesAndUpdates(lines)
		third := updates[2]
		sorted := fixSorting(third, rules)
		expected := []int{97, 75, 47, 29, 13}

		if !slices.Equal(sorted, expected) {
			t.Errorf("not equal")
			fmt.Println("sorted", sorted)
			fmt.Println("expected", expected)
		}
	})
}

func TestRuleExistsByXY(t *testing.T) {
	t.Run("Can get something", func(t *testing.T) {
		rules := make([]Rule, 0)
		rules = append(rules, Rule{x: 0, y: 1})
		rules = append(rules, Rule{x: 1, y: 0})
		rules = append(rules, Rule{x: 1, y: 1})
		rules = append(rules, Rule{x: 0, y: 0})

		if !hasRuleForXY(0, 1, rules) {
			t.Errorf("Couldnt find 0, 1")
		}

		if !hasRuleForXY(1, 0, rules) {
			t.Errorf("Couldnt find 1, 0")
		}

		if !hasRuleForXY(1, 1, rules) {
			t.Errorf("Couldnt find 1, 1")
		}

		if !hasRuleForXY(0, 0, rules) {
			t.Errorf("Couldnt find 0, 0")
		}

		if hasRuleForXY(2, 2, rules) {
			t.Errorf("Should be impossible")
		}
	})
}

func TestCountingMiddlesPartTwo(t *testing.T) {
	t.Run("Can get total amount of all correct middles", func(t *testing.T) {
		lines := createTestInputPartTwo()
		rules, updates := separateRulesAndUpdates(lines)
		result := checkUpdatesAgainstRulesAndAttemptSort(rules, updates)

		total := 0
		expected := 123
		for _, count := range result {
			total += count
		}

		if total != expected {
			t.Errorf("got %d, want %d", total, expected)
		}

	})
}
