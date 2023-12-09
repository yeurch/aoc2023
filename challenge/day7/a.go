package day7

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yeurch/aoc2023/challenge"
)

type HandType int

const (
	HIGH_CARD HandType = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 7, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	return doImpl(challenge, false)
}

func getCardValueA(c byte) int {
	if c >= '2' && c <= '9' {
		val, _ := strconv.Atoi(string(c))
		return val - 2 // 2 is lowest, with score of 0.
	}
	switch c {
	case 'T':
		return 8
	case 'J':
		return 9
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

func getCategoryA(hand string) HandType {
	if len(hand) != 5 {
		panic("Hand must contain five cards.")
	}
	cardCounts := make(map[rune]int)
	for _, c := range hand {
		cardCounts[c]++
	}

	counts := make([]int, 0, len(cardCounts))
	for _, v := range cardCounts {
		counts = append(counts, v)
	}
	sort.Ints(counts)

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
