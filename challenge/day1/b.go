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
	digits := map[string]rune{
		"one": '1', "two": '2', "three": '3', "four": '4', "five": '5',
		"six": '6', "seven": '7', "eight": '8', "nine": '9',
	}
	for line := range challenge.Lines() {
		s := make([]rune, 0)
		for i, c := range line {
			if unicode.IsNumber(c) {
				s = append(s, c)
				continue
			}
			for k, v := range digits {
				if strings.HasPrefix(string(line[i:]), k) {
					s = append(s, v)
				}
			}
		}
		first := int(s[0] - '0')
		last := int(s[len(s)-1] - '0')
		result += 10*first + last
	}
	return result
}
