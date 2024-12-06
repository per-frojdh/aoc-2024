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

func printMap(grid [][]rune) {
	for row := range grid {
		for col := range grid[row] {
			fmt.Printf("%c", grid[row][col])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("----------\n")
}

func TestParseMapToGrid(t *testing.T) {
	t.Run("Can parse example input into grid", func(t *testing.T) {
		lines := createTestinput()
		grid, _ := parseMapToGrid(lines)

		expectedLength := 10
		expectedDepth := 10

		// printMap()

		if len(grid) != expectedLength {
			t.Errorf("got %d, want %d", len(grid), expectedLength)
		}

		for i := range grid {
			if len(grid[i]) != expectedDepth {
				t.Errorf("got %d, want %d at index %d", len(grid[i]), expectedDepth, i)
			}
		}
	})

	t.Run("Can find start position from example grid", func(t *testing.T) {
		lines := createTestinput()
		_, start := parseMapToGrid(lines)

		expectedStart := []int{6, 4}

		if !slices.Equal(start, expectedStart) {
			t.Errorf("got %d, want %d", start, expectedStart)
		}
	})
}

func TestCanMoveDirection(t *testing.T) {
	t.Run("Can move into a non-visited space from example start", func(t *testing.T) {
		lines := createTestinput()
		grid, start := parseMapToGrid(lines)

		expected_pos := []int{5, 4}
		expected_block := '.'
		allowed, block, new_pos := canMoveDirection(grid, start, "Up")

		if !allowed {
			t.Errorf("got %t, want %t", allowed, true)
		}

		if !slices.Equal(new_pos, expected_pos) {
			t.Errorf("got %d, want %d", new_pos, expected_pos)
		}

		if block != expected_block {
			t.Errorf("got %c, want %c", block, expected_block)
		}
	})

	t.Run("Can not move into an out of bounds space", func(t *testing.T) {
		lines := createTestinput()
		grid, _ := parseMapToGrid(lines)

		expected_pos := []int{0, 0}
		custom_start := []int{0, 0}
		expected_block := '_'
		finished, block, new_pos := canMoveDirection(grid, custom_start, "Up")

		if finished {
			t.Errorf("got %t, want %t", finished, false)
		}

		if !slices.Equal(new_pos, expected_pos) {
			t.Errorf("got %d, want %d", new_pos, expected_pos)
		}

		if block != expected_block {
			t.Errorf("got %c, want %c", block, expected_block)
		}
	})

	t.Run("Can not move into a blocked space", func(t *testing.T) {
		lines := createTestinput()
		grid, _ := parseMapToGrid(lines)

		expected_pos := []int{0, 3}
		custom_start := []int{0, 3}
		expected_block := '#'
		_, block, new_pos := canMoveDirection(grid, custom_start, "Right")

		if !slices.Equal(new_pos, expected_pos) {
			t.Errorf("got %d, want %d", new_pos, expected_pos)
		}

		if block != expected_block {
			t.Errorf("got %c, want %c", block, expected_block)
		}
	})

	t.Run("Can not move down into a blocked space", func(t *testing.T) {
		lines := createTestinput()
		grid, _ := parseMapToGrid(lines)

		expected_pos := []int{0, 9}
		custom_start := []int{0, 9}
		expected_block := '#'
		_, block, new_pos := canMoveDirection(grid, custom_start, "Down")

		if !slices.Equal(new_pos, expected_pos) {
			t.Errorf("got %d, want %d", new_pos, expected_pos)
		}

		if block != expected_block {
			t.Errorf("got %c, want %c", block, expected_block)
		}
	})

	t.Run("Can start moving through the example", func(t *testing.T) {
		lines := createTestinput()
		grid, start := parseMapToGrid(lines)

		expected_visits := 41
		visits := startMoving(grid, start, func(i int, g []int) {
			// time.Sleep(1 * time.Second)
			// printMap(grid)
		})

		if visits != expected_visits {
			t.Errorf("got %d, want %d", visits, expected_visits)
		}
	})
}

func TestGetNextDirection(t *testing.T) {
	t.Run("Can get next direction from up", func(t *testing.T) {
		direction := "Up"
		expected_direction := "Right"
		got := getNextDirection(direction)

		if got != expected_direction {
			t.Errorf("got %q, want %q", got, expected_direction)
		}
	})

	t.Run("Can get next direction from right", func(t *testing.T) {
		direction := "Right"
		expected_direction := "Down"
		got := getNextDirection(direction)

		if got != expected_direction {
			t.Errorf("got %q, want %q", got, expected_direction)
		}
	})

	t.Run("Can get next direction from Down", func(t *testing.T) {
		direction := "Down"
		expected_direction := "Left"
		got := getNextDirection(direction)

		if got != expected_direction {
			t.Errorf("got %q, want %q", got, expected_direction)
		}
	})

	t.Run("Can get next direction from left", func(t *testing.T) {
		direction := "Left"
		expected_direction := "Up"
		got := getNextDirection(direction)

		if got != expected_direction {
			t.Errorf("got %q, want %q", got, expected_direction)
		}
	})
}

func TestCanUpdateGrid(t *testing.T) {
	t.Run("Can update grid with a new rune", func(t *testing.T) {
		lines := createTestinput()
		grid, start := parseMapToGrid(lines)

		grid = updateGrid(grid, start, '#')
		expected_block := '#'

		if grid[start[0]][start[1]] != expected_block {
			t.Errorf("got %c, want %c", grid[start[0]][start[1]], expected_block)
		}
	})
}

func TestCanGetWindow(t *testing.T) {
	_, lines := util.ReadInputIntoLines("test_window.txt")
	grid, _ := parseMapToGrid(lines)

	window := getWindow(grid, []int{4, 5}, 3)
	fmt.Println(window)
}
