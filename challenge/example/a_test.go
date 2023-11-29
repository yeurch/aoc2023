package example

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yeurch/aoc2023/challenge"
)

func TestA(t *testing.T) {
	input := challenge.FromLiteral("42")

	result := a(input)

	require.Equal(t, 42, result)
}
