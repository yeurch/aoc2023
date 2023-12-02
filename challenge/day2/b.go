package day2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yeurch/aoc2023/challenge"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 2, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	re := regexp.MustCompile(`Game (\d+): (.+)`)
	result := 0 // This will store the sum of indexes

	for line := range challenge.Lines() {
		m := re.FindStringSubmatch(line)
		grabs := strings.Split(m[2], "; ")
		maxRed := 0
		maxGreen := 0
		maxBlue := 0
		for _, grab := range grabs {
			ballSets := strings.Split(grab, ", ")
			for _, ballSet := range ballSets {
				flds := strings.Fields(ballSet)
				num, _ := strconv.Atoi(flds[0])
				color := flds[1]
				if color == "red" && num > maxRed {
					maxRed = num
				}
				if color == "green" && num > maxGreen {
					maxGreen = num
				}
				if color == "blue" && num > maxBlue {
					maxBlue = num
				}
			}
		}
		result += maxRed * maxGreen * maxBlue
	}
	return result
}
