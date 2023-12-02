package day2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yeurch/aoc2023/challenge"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 2, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	re := regexp.MustCompile(`Game (\d+): (.+)`)
	result := 0 // This will store the sum of indexes

	for line := range challenge.Lines() {
		m := re.FindStringSubmatch(line)
		idx, _ := strconv.Atoi(m[1])
		grabs := strings.Split(m[2], "; ")
		ok := true
		for _, grab := range grabs {
			ballSets := strings.Split(grab, ", ")
			for _, ballSet := range ballSets {
				flds := strings.Fields(ballSet)
				num, _ := strconv.Atoi(flds[0])
				color := flds[1]
				if (color == "red" && num > 12) || (color == "green" && num > 13) ||
					(color == "blue" && num > 14) {
					ok = false
				}
			}
		}
		if ok {
			result += idx
		}
	}
	return result
}
