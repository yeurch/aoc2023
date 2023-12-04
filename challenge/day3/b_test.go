package day3

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yeurch/aoc2023/challenge"
)

func TestB(t *testing.T) {
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

	result := partB(input)

	require.Equal(t, 467835, result)
}
