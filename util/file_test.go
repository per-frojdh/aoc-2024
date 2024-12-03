package util

import (
	"testing"
)

func TestFileReading(t *testing.T) {
	t.Run("Can read a input file without a final newline", func(t *testing.T) {
		_, lines := ReadInputIntoLines("test-files/input.txt")

		expected := 6

		if len(lines) != expected {
			t.Errorf("got %d, want %d", len(lines), expected)
		}
	})

	t.Run("Can read a input file with a final newline, but not count it", func(t *testing.T) {
		_, lines := ReadInputIntoLines("test-files/input-with-empty-eol.txt")

		expected := 6

		if len(lines) != expected {
			t.Errorf("got %d, want %d", len(lines), expected)
		}
	})

	t.Run("Can read an input file into a single string", func(t *testing.T) {
		_, line := ReadInputIntoString("test-files/input-single-string.txt")

		expected := "ABC DE F"

		if line != expected {
			t.Errorf("got %q, want %q", line, expected)
		}

	})

}
