package day7

import (
	"sort"
	"strconv"
	"strings"

	"github.com/yeurch/aoc2023/challenge"
)

type Hand struct {
	cards string
	bid   int
	score int
}

func doImpl(challenge *challenge.Input, isPartB bool) int {
	lines := challenge.LineSlice()
	hands := make([]Hand, len(lines))
	for i, line := range lines {
		items := strings.Fields(line)
		cards := items[0]
		bid, _ := strconv.Atoi(items[1])
		hands[i] = Hand{
			cards, bid, getHandScore(cards, isPartB),
		}
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].score < hands[j].score
	})

	result := 0
	for i, h := range hands {
		result += (i + 1) * h.bid
	}
	return result
}

func getHandScore(hand string, isPartB bool) int {
	mult := 1
	result := 0
	for i := 4; i >= 0; i-- {
		result += mult * getCardValue(hand[i], isPartB)
		mult *= 13
	}
	result += mult * int(getCategory(hand, isPartB))
	return result
}

func getCardValue(c byte, isPartB bool) int {
	if isPartB {
		return getCardValueB(c)
	}
	return getCardValueA(c)
}

func getCategory(hand string, isPartB bool) HandType {
	if isPartB {
		return getCategoryB(hand)
	}
	return getCategoryA(hand)
}
