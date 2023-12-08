package day5

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yeurch/aoc2023/challenge"
	"github.com/yeurch/aoc2023/util"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 5, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	lines := challenge.LineSlice()

	seedNums := util.NewSet[int]()
	seedNumsStr := strings.Fields(strings.Split(lines[0], ":")[1])
	for _, s := range seedNumsStr {
		val, _ := strconv.Atoi(s)
		seedNums.Add(val)
	}

	return doImpl(seedNums, lines[1:])
}
