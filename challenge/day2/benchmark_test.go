package day2

import (
	"testing"

	"github.com/yeurch/aoc2023/challenge"
)

func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = partA(challenge.FromFile())
	}
}

func BenchmarkB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = partB(challenge.FromFile())
	}
}
