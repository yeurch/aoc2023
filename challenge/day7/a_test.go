package day7

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yeurch/aoc2023/challenge"
)

func TestGetHandType(t *testing.T) {
	var input string
	var result HandType

	input = "35355"
	result = getCategory(input, false)
	require.Equal(t, FULL_HOUSE, result)
}

func TestA(t *testing.T) {
	input := challenge.FromLiteral(`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`)

	result := partA(input)

	require.Equal(t, 6440, result)
}
