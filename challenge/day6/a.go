package day6

import (
	"fmt"
	"math"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yeurch/aoc2023/challenge"
	"github.com/yeurch/aoc2023/util"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 6, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

/*
* Let:
*   d = length to beat the record by /mm
*   R = record distance
*   T = length of race /ms
*   s = speed
*   tc = time charging
*   ts = time sailing
*
*   distance = speed x time gives us
*   d = s.T - R
*   But s = tc; and T = tc + ts, so:
*   d = tc.(T-tc) - R
*   d = (-tc)^2 +T.tc - R
*
*   As this is a negative squared formula, we expect a single maximum, and we have to find
*   the extend at the points at which d is positive.  So solving for roots gives:
*   a = -1
*   b = T
*   c = -R
*
*   (-b +- (b^2 - 4.a.c)^0.5)/(2.a)
*   (-T +- (T^2 - 4.R)^0.5)/(-2)
*   (T +- (T^2 - 4.R)^0.5)/2
 */
func partA(challenge *challenge.Input) int {
	lines := challenge.LineSlice()
	times, _ := util.ParseIntSlice(strings.Split(lines[0], ":")[1])
	records, _ := util.ParseIntSlice(strings.Split(lines[1], ":")[1])

	result := 1
	for i := 0; i < len(times); i++ {
		time := float64(times[i])
		record := float64(records[i])

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
		result *= int(math.Floor(maxVal)) - int(math.Ceil(minVal)) + 1 - excludedLimits
	}
	return result
}
