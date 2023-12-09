package day7

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yeurch/aoc2023/challenge"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 7, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

func partB(challenge *challenge.Input) int {
	return doImpl(challenge, true)
}

func getCardValueB(c byte) int {
	if c >= '2' && c <= '9' {
		val, _ := strconv.Atoi(string(c))
		return val - 1 // 2 is lowest, with score of 1 (because J is lowest).
	}
	switch c {
	case 'T':
		return 9
	case 'J':
		return 0
	case 'Q':
		return 10
	case 'K':
		return 11
	case 'A':
		return 12
	default:
		panic("Unexpected card value")
	}
}

func getCategoryB(hand string) HandType {
	if len(hand) != 5 {
		panic("Hand must contain five cards.")
	}
	cardCounts := make(map[rune]int)
	for _, c := range hand {
		cardCounts[c]++
	}

	numJacks := 0
	counts := make([]int, 0, len(cardCounts))
	for k, v := range cardCounts {
		if k == 'J' {
			numJacks += v
			continue
		}
		counts = append(counts, v)
	}
	sort.Ints(counts)

	// Our Jacks should all be classed as the card we already have the most of.
	if len(counts) == 0 {
		counts = []int{0}
	}
	counts[len(counts)-1] += numJacks

	var sb strings.Builder
	for _, c := range counts {
		sb.WriteString(fmt.Sprintf("%v", c))
	}
	handTypeStr := sb.String()

	switch handTypeStr {
	case "5":
		return FIVE_OF_A_KIND
	case "14":
		return FOUR_OF_A_KIND
	case "23":
		return FULL_HOUSE
	case "113":
		return THREE_OF_A_KIND
	case "122":
		return TWO_PAIR
	case "1112":
		return ONE_PAIR
	default:
		return HIGH_CARD
	}
}
