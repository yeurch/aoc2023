package day3

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yeurch/aoc2023/challenge"
)

func TestA(t *testing.T) {
	input := challenge.FromLiteral(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`)

	result := partA(input)

	require.Equal(t, 4361, result)
}

func TestFindExtent(t *testing.T) {
	input := "$8..334"

	result, start, length := findExtent(input, 0)
	require.Equal(t, 8, result)
	require.Equal(t, 1, start)
	require.Equal(t, 1, length)

	result, start, length = findExtent(input, start+length+1)
	require.Equal(t, 334, result)
	require.Equal(t, 4, start)
	require.Equal(t, 3, length)

	result, _, _ = findExtent(input, start+length+1)
	require.Less(t, result, 0)
}

func TestFindExtent2(t *testing.T) {
	input := "8.42."

	result, start, length := findExtent(input, 0)
	require.Equal(t, 8, result)
	require.Equal(t, 0, start)
	require.Equal(t, 1, length)

	result, start, length = findExtent(input, start+length+1)
	require.Equal(t, 42, result)
	require.Equal(t, 2, start)
	require.Equal(t, 2, length)

	result, _, _ = findExtent(input, start+length+1)
	require.Less(t, result, 0)
}

func TestFindExtent3(t *testing.T) {
	input := "12345"

	result, start, length := findExtent(input, 0)
	require.Equal(t, 12345, result)
	require.Equal(t, 0, start)
	require.Equal(t, 5, length)

	result, _, _ = findExtent(input, start+length+1)
	require.Less(t, result, 0)
}
