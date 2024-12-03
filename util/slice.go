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
