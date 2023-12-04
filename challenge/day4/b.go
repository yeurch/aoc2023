package day4

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yeurch/aoc2023/challenge"
	"github.com/yeurch/aoc2023/util"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 4, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	re := regexp.MustCompile(`Card +\d+: ([\d ]+) \| ([\d ]+)$`)

	result := 0
	lines := challenge.LineSlice()
	numOrigCards := len(lines)
	copies := make([]int, numOrigCards)

	for i, line := range lines {
		copies[i]++
		m := re.FindStringSubmatch(line)

		ourNumbersStr := strings.Fields(m[1])
		ourNumbers := util.NewSet[int]()
		for _, s := range ourNumbersStr {
			val, _ := strconv.Atoi(s)
			ourNumbers.Add(val)
		}

		theirNumbersStr := strings.Fields(m[2])
		theirNumbers := util.NewSet[int]()
		for _, s := range theirNumbersStr {
			val, _ := strconv.Atoi(s)
			theirNumbers.Add(val)
		}

		numWinningNumbers := ourNumbers.Intersect(theirNumbers).Cardinality()
		for j := i + 1; j <= i+numWinningNumbers; j++ {
			if j < numOrigCards {
				copies[j] += copies[i]
			}
		}
	}

	for _, i := range copies {
		result += i
	}
	return result
}
