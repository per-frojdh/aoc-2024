package main

import (
	"aoc/util"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func separateRulesAndUpdates(lines []string) ([]Rule, [][]int) {
	rules := []Rule{}
	updates := make([][]int, 0)

	for i := range lines {
		current := lines[i]

		if strings.Contains(current, "|") {
			rule := strings.Split(current, "|")

			n1, n1_err := strconv.Atoi(rule[0])
			n2, n2_err := strconv.Atoi(rule[1])

			if n1_err != nil || n2_err != nil {
				fmt.Printf("Failed to parse rule\n")
				fmt.Println(n1, n2)
				continue
			}

			rules = append(rules, Rule{x: n1, y: n2})
			continue
		}

		if strings.Contains(current, ",") {
			update := strings.Split(current, ",")
			arr := []int{}

			for y := range update {
				n, err := strconv.Atoi(update[y])
				if err != nil {
					fmt.Printf("Failed to parse update")
					fmt.Println(n)
					continue
				}

				arr = append(arr, n)

			}
			updates = append(updates, arr)
			continue
		}
	}

	return rules, updates
}

type Rule struct {
	x, y int
}

func checkCorrectness(lines []int, rules []Rule) bool {
	// Iterate over every rule we have collected
	for r := range rules {
		rule := rules[r]

		// Set initial positions
		index_of_x := -1
		index_of_y := -1

		// Loop over the line of updates and check each value
		// against the "definition" of the rule (x) and
		// the allowed values (y)
		for pos, update := range lines {
			if update == rule.x {
				index_of_x = pos
			}

			if update == rule.y {
				index_of_y = pos
			}
		}

		// If we found both x and y, and x is found later than y
		// it's invalid
		if index_of_x != -1 && index_of_y != -1 && index_of_x > index_of_y {
			return false
		}
	}
	return true
}

func checkUpdatesAgainstRules(rules []Rule, updates [][]int) map[int]int {
	updateResult := map[int]int{}
	for i, update := range updates {
		// For each update, check the correctness according to all rules
		if is_correct := checkCorrectness(update, rules); is_correct {
			// If it's correct, find the middle number in the update line
			// and add it to the total
			updateResult[i] = update[len(update)/2]
		}
	}

	return updateResult
}

func checkUpdatesAgainstRulesAndAttemptSort(rules []Rule, updates [][]int) map[int]int {
	updateResult := map[int]int{}
	for i, update := range updates {
		// For each update, check the correctness according to all rules
		if is_correct := checkCorrectness(update, rules); !is_correct {
			// If it's incorrect, attempt to sort
			sorted := fixSorting(update, rules)
			updateResult[i] = sorted[len(sorted)/2]
		}
	}

	return updateResult

}

func hasRuleForXY(x, y int, rules []Rule) bool {
	for _, rule := range rules {
		if rule.x == x && rule.y == y {
			return true
		}
	}

	return false
}

func fixSorting(line []int, rules []Rule) []int {
	sortedLine := slices.Clone(line)

	// Check rules for both directions, both from x -> y and
	// from y -> x.
	slices.SortFunc(sortedLine, func(x, y int) int {

		// If there's a rule for y to be before x, don't sort it
		// y should remain after x
		if hasRuleForXY(y, x, rules) {
			return 1
		}

		// If there's a rule for x to be before y, sort it.
		if hasRuleForXY(x, y, rules) {
			return -1
		}

		// Default comparison
		return x - y
	})

	return sortedLine
}

func main() {
	success, lines := util.ReadInputIntoLines("input.txt")
	if success {
		rules, updates := separateRulesAndUpdates(lines)
		result := checkUpdatesAgainstRules(rules, updates)
		total := 0
		for _, count := range result {
			total += count
		}

		fmt.Printf("Part1: Sum of all middles are %d\n", total)
		result = checkUpdatesAgainstRulesAndAttemptSort(rules, updates)
		total = 0
		for _, count := range result {
			total += count
		}

		fmt.Printf("Part2: Sum of all middles are after sorting %d", total)

	}
}
