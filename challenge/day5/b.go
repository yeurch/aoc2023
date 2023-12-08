package day5

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yeurch/aoc2023/challenge"
	"github.com/yeurch/aoc2023/util"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 5, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	lines := challenge.LineSlice()

	almanac := parseAlmanac(lines[1:])

	seamValues := util.NewSet[int]() // We'll store the seed value for each time we break 1:1 linear growth

	// Numbers in each mapping are [ dest_start_range, source_start_range, num_values ]
	// Our first seams will be the source_start_range values of the final mapping
	for _, rule := range almanac[len(almanac)-1] {
		seamValues.Add(rule[1])
	}

	// We need to start with the last-but-one set of mappings, and work forwards to the seed ids.
	// Each set of mappings will add more seam values to our list.

	// The values in the set are the source_start_range of the new set, plus the mappings into the source_start_range
	// of everything that's already in the set.
	for i := len(almanac) - 2; i >= 0; i-- {
		newSeams := util.NewSet[int]()
		mapping := almanac[i]
		for _, rule := range mapping {
			newSeams.Add(rule[1])
		}
		for val := range seamValues.Iter() {
			newSeam := reverseLookup(val, mapping)
			newSeams.Add(newSeam)
		}
		seamValues = newSeams
	}

	completeMapping := make(map[int]int)
	for seamValue := range seamValues.Iter() {
		completeMapping[seamValue] = lookup(seamValue, almanac)
	}

	seamValuesSlice := make([]int, 0, seamValues.Cardinality())
	for val := range seamValues.Iter() {
		seamValuesSlice = append(seamValuesSlice, val)
	}
	sort.Ints(seamValuesSlice)

	// Now, let's actually look at our input and figure out the answer
	seedNumsStr := strings.Fields(strings.Split(lines[0], ":")[1])

	result := math.MaxInt
	for i := 0; i < len(seedNumsStr); i += 2 {
		start, _ := strconv.Atoi(seedNumsStr[i])
		length, _ := strconv.Atoi(seedNumsStr[i+1])

		loc := lookup(start, almanac)
		if loc < result {
			result = loc
		}
		for _, seamVal := range seamValuesSlice {
			if seamVal > start && seamVal < start+length {
				// We have another possible min value, as we're at a seam in our
				// almanac within the range of seed values.
				loc = lookup(seamVal, almanac)
				if loc < result {
					result = loc
				}
			}
		}
	}

	return result
}
