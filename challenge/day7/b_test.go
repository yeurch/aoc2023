package day7

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yeurch/aoc2023/challenge"
)

func TestB(t *testing.T) {
	input := challenge.FromLiteral(`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`)

	result := partB(input)

	require.Equal(t, 5905, result)
}
