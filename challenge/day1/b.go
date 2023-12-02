package day1

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
	"github.com/yeurch/aoc2023/challenge"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 1, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	result := 0
	digits := []string{
		"one", "two", "three", "four", "five",
		"six", "seven", "eight", "nine",
	}
	for line := range challenge.Lines() {
		s := make([]int, 0)
		for i, c := range line {
			if unicode.IsNumber(c) {
				s = append(s, int(c-'0'))
				continue
			}
			for j, digit := range digits {
				if strings.HasPrefix(string(line[i:]), digit) {
					s = append(s, j+1)
				}
			}
		}
		first := s[0]
		last := s[len(s)-1]
		result += 10*first + last
	}
	return result
}
