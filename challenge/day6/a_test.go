package day6

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yeurch/aoc2023/challenge"
)

func TestA(t *testing.T) {
	input := challenge.FromLiteral(`Time:      7  15   30
Distance:  9  40  200`)

	result := partA(input)

	require.Equal(t, 288, result)
}
