package util

import (
	"testing"
)

func TestCanUseAt(t *testing.T) {
	t.Run("Can use At to find a valid position", func(t *testing.T) {
		arr := []int{1, 2, 3, 4}
		expected := 3
		got, err := At(arr, 2)

		if err != nil {
			t.Error("should not get an error from a valid position")
		}

		if expected != got {
			t.Errorf("got %d, want %d", got, expected)
		}
	})

	t.Run("Can use At to find a invalid position", func(t *testing.T) {
		arr := []int{1, 2, 3, 4}
		_, err := At(arr, 10)

		if err == nil {
			t.Error("should get an error from a invalid position")
		}
	})
}
