package main

import (
	"aoc/util"
	"fmt"
	"strings"
)

var directionMap = map[string][]int{
	"Up":    {-1, 0},
	"Right": {0, 1},
	"Down":  {1, 0},
	"Left":  {0, -1},
}

var blockMap = map[string]rune{
	"Available": '.',
	"Blocked":   '#',
	"Visited":   'X',
	"Undefined": '_',
}

func getGuardRuneByDirection(direction string) rune {
	switch direction {
	case "Up":
		return '^'
	case "Right":
		return '>'
	case "Down":
		return 'v'
	case "Left":
		return '<'
	default:
		panic("Invalid direction supplied")
	}
}

func getNextDirection(current string) string {
	// Define the order
	order := []string{"Up", "Right", "Down", "Left"}

	// Find current index
	currentIndex := -1
	for i, dir := range order {
		if dir == current {
			currentIndex = i
			break
		}
	}

	// Return next direction (wrapping around to start if at end)
	nextIndex := (currentIndex + 1) % len(order)
	return order[nextIndex]
}

func getDirection(direction string) []int {
	return directionMap[direction]
}

func updateGrid(grid [][]rune, pos []int, r rune) [][]rune {
	_, exists := util.GridAt(grid, pos[0], pos[1])

	if exists == nil {

		grid[pos[0]][pos[1]] = r
	}
	return grid
}

func parseMapToGrid(lines []string) ([][]rune, []int) {
	start := []int{}
	grid := [][]rune{}

	for i := range lines {
		if match := strings.Index(lines[i], "^"); match != -1 {
			start = []int{i, match}
		}

		grid = append(grid, []rune(lines[i]))
	}
	return grid, start
}

func canMoveDirection(grid [][]rune, start []int, direction string) (bool, rune, []int) {
	dir := getDirection(direction)
	next_pos := []int{start[0] + dir[0], start[1] + dir[1]}

	content, within_grid := util.GridAt(grid, next_pos[0], next_pos[1])
	if within_grid != nil {
		fmt.Println("Should be finishing now")
		return false, blockMap["Undefined"], start
	}

	if content == blockMap["Blocked"] {
		fmt.Println("Blocked, rotating direction")
		return true, blockMap["Blocked"], start
	}

	return true, content, next_pos
}

type moving_callback func(int, []int)

func startMoving(grid [][]rune, start []int, cb moving_callback) int {
	visits := 0
	direction := "Up"
	guard_position := start

	for {
		cb(visits, guard_position)
		if finished, block, position := canMoveDirection(grid, guard_position, direction); finished {
			if block == blockMap["Blocked"] {
				// If we are blocked, we get a new direction
				// but we don't move the guard
				direction = getNextDirection(direction)
				// Rotate the guard inside the grid
				grid = updateGrid(grid, guard_position, getGuardRuneByDirection(direction))
				continue
			}

			if block == blockMap["Available"] {
				// We are free to move into the next one
				// update guard_position and add +1 to visits
				// Update the available position to now be visited
				grid = updateGrid(grid, guard_position, blockMap["Visited"])
				guard_position = position

				// Move the guard to the new position
				grid = updateGrid(grid, position, getGuardRuneByDirection(direction))
				visits++
				continue
			}

			if block == blockMap["Visited"] {
				grid = updateGrid(grid, guard_position, blockMap["Visited"])
				guard_position = position
				// // Update the new guard position
				grid = updateGrid(grid, guard_position, getGuardRuneByDirection(direction))
				continue
			}
		} else {
			// We are about to move out of bounds, which means the
			// patrol has finished, increment visits and remove the
			// guard from the grid
			fmt.Println("Should be finished now?")
			visits++
			grid = updateGrid(grid, guard_position, blockMap["Visited"])
			cb(visits, guard_position)
			return visits
		}
	}
}

func getWindow(grid [][]rune, guard []int, offset int) [][]rune {
	gx := guard[0]
	// gy := guard[1]
	grid_length := len(grid)
	// grid_depth := len(grid[0])

	x_offset_start := gx - offset
	x_offset_end := gx + offset

	if x_offset_start < 0 {
		x_offset_start = 0
	}

	if x_offset_end > grid_length {
		x_offset_end = grid_length
	}

	fmt.Println(gx)
	fmt.Println(x_offset_start)
	fmt.Println(x_offset_end)
	// window_x := []int{gx - x_offset, gx + x_offset}
	window_x := grid[x_offset_start:x_offset_end]
	// window_y := window_x[gy-y_offset : gy+y_offset]

	// for _, row := range window_x {
	// 	for _, col := range window_y {
	// 		fmt.Printf("%c", row[col])
	// 	}
	// 	fmt.Printf("\n")
	// }
	// fmt.Printf("----------\n")
	return window_x
}

func main() {
	success, lines := util.ReadInputIntoLines("input.txt")
	if success {
		visits := 0
		grid, start := parseMapToGrid(lines)
		visits = startMoving(grid, start, func(i int, guard []int) {
			// if i == 3462 {
			// 	printAroundGuard(grid)
			// 	panic("End me now")
			// }
			fmt.Println(i)
		})
		fmt.Printf("Part1: Amount of visits is %d\n", visits)
	}
}
