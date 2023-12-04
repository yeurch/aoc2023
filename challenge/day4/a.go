package day4

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yeurch/aoc2023/challenge"
	"github.com/yeurch/aoc2023/util"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 4, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	re := regexp.MustCompile(`Card +\d+: ([\d ]+) \| ([\d ]+)$`)

	result := 0
	for line := range challenge.Lines() {
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

		winningNumbers := ourNumbers.Intersect(theirNumbers)
		result += int(math.Pow(2, float64(winningNumbers.Cardinality()-1)))
	}
	return result
}
