package day8

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yeurch/aoc2023/challenge"
)

func TestA(t *testing.T) {
	t.Skipf("Challenge not yet solved")
	input := challenge.FromLiteral("foobar")

	result := partA(input)

	require.Equal(t, 42, result)
}
