package util

import (
	"strconv"
	"strings"
)

// Reads a line of numbers, seperated by one or more whitespace characters.
// Returns an integer slice containing the parsed numbers, and an error if appropriate.
func ParseIntSlice(s string) ([]int, error) {
	fields := strings.Fields(s)
	result := make([]int, len(fields))
	for i, strVal := range fields {
		val, err := strconv.Atoi(strVal)
		if err != nil {
			return nil, err
		}
		result[i] = val
	}
	return result, nil
}
