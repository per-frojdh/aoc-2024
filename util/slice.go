package util

import "fmt"

// https://www.reddit.com/r/golang/comments/183ymo1/comment/karz6sx/
func At[T any](ar []T, i int) (T, error) {
	if i >= 0 && i < len(ar) {
		return ar[i], nil
	}
	var z T
	return z, fmt.Errorf("out of bounds")
}

// This is just an adjusted version of the At method above
// but since I don't know generics well enough, this one handles
// 2d arrays
func GridAt[T any](ar [][]T, i, y int) (T, error) {
	if i >= 0 && i < len(ar) {
		if y >= 0 && y < len(ar[i]) {
			return ar[i][y], nil
		}
	}
	var z T
	return z, fmt.Errorf("out of bounds")
}
