package day6

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yeurch/aoc2023/challenge"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 6, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	lines := challenge.LineSlice()

	timeStr, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(lines[0], ":")[1], " ", ""))
	time := float64(timeStr)
	recordStr, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(lines[1], ":")[1], " ", ""))
	record := float64(recordStr)

	plusMinusTerm := math.Sqrt(math.Pow(time, 2.0) - 4.0*record)
	minVal := (time - plusMinusTerm) / 2
	maxVal := (time + plusMinusTerm) / 2

	excludedLimits := 0
	if math.Ceil(minVal) == minVal {
		excludedLimits++
	}
	if math.Floor(maxVal) == maxVal {
		excludedLimits++
	}
	result := int(math.Floor(maxVal)) - int(math.Ceil(minVal)) + 1 - excludedLimits
	return result
}
