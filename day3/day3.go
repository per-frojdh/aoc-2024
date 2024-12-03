package main

import (
	"aoc/util"
	"fmt"
	"regexp"
	"strconv"
)

const (
	Active   = 1
	NoChange = 0
	Inactive = -1
)

func parseCorruptedWithExtraInstructions(input string) [][]int {
	r, _ := regexp.Compile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	do := Active

	result := [][]int{}
	matches := r.FindAllStringSubmatch(input, -1)

	handleInstruction := func(str []string) int {
		if str[0] == "do()" {
			return Active
		}

		if str[0] == "don't()" {
			return Inactive
		}

		return NoChange
	}

	for _, n := range matches {

		// An index of a submatch will never contain both
		// a "do/dont" command, an a mul-instruction, so if we find one
		// change the state and move to the next submatch.
		instruction := handleInstruction(n)
		if instruction != NoChange {
			do = instruction
			continue
		}

		if do == Active {
			n1, err := strconv.Atoi(n[1])
			if err != nil {
				fmt.Printf("Could not convert left-hand values from regex: %v \n", n)
				continue
			}

			n2, err := strconv.Atoi(n[2])
			if err != nil {
				fmt.Printf("Could not convert right-hand values from regex: %v \n", n)
				continue
			}
			result = append(result, []int{n1, n2})
		}
	}

	return result
}

func parseCorruptedInput(input string) (bool, [][]int) {
	r, _ := regexp.Compile(`mul\((\d+),(\d+)\)`)

	result := [][]int{}

	matches := r.FindAllStringSubmatch(input, -1)
	for _, n := range matches {
		n1, err := strconv.Atoi(n[1])
		if err != nil {
			fmt.Printf("Could not convert left-hand values from regex: %v \n", n)
			continue
		}

		n2, err := strconv.Atoi(n[2])
		if err != nil {
			fmt.Printf("Could not convert right-hand values from regex: %v \n", n)
			continue
		}
		result = append(result, []int{n1, n2})
	}

	return len(result) != 0, result
}

func computeCorruptedInput(input [][]int) int {
	result := 0
	for i := range input {
		first, err := util.At(input[i], 0)
		if err != nil {
			fmt.Printf("Could not get left-hand values from slice: %v \n", input[i])
			continue
		}

		second, err := util.At(input[i], 1)
		if err != nil {
			fmt.Printf("Could not get right-hand values from slice: %v \n", input[i])
			continue
		}

		result += first * second
	}

	return result
}

func main() {
	success, lines := util.ReadInputIntoString("input.txt")
	if success {
		_, parsed := parseCorruptedInput(lines)
		result := computeCorruptedInput(parsed)

		fmt.Printf("Part1: Result from all sets of computation is %d \n", result)

		parsed = parseCorruptedWithExtraInstructions(lines)
		result = computeCorruptedInput(parsed)

		fmt.Printf("Part2: Result from all sets of computation with extra instructions is %d \n", result)
	}
}
