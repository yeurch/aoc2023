package day1

import (
	"fmt"
	"unicode"

	"github.com/spf13/cobra"
	"github.com/yeurch/aoc2023/challenge"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 1, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	result := 0
	for line := range challenge.Lines() {
		s := make([]rune, 0)
		for _, c := range line {
			if unicode.IsNumber(c) {
				s = append(s, c)
			}
		}
		first := int(s[0] - '0')
		last := int(s[len(s)-1] - '0')
		result += 10*first + last
	}
	return result
}
