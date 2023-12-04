package day3

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/spf13/cobra"
	"github.com/yeurch/aoc2023/challenge"
)

func aCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "a",
		Short: "Day 3, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partA(challenge.FromFile()))
		},
	}
}

func partA(challenge *challenge.Input) int {
	result := 0
	grid := challenge.TileMap()
	_, height := grid.Size()
	for ln := 0; ln < height; ln++ {
		s := grid.LineStr(ln)

		pos := 0
		for {
			val, start, length := findExtent(s, pos)
			pos = start + length + 1
			if val < 0 {
				break
			}

			completedValue := false
			for y := ln - 1; y < ln+2; y++ {
				for x := start - 1; x < start+length+1; x++ {
					r, ok := grid.TileAt(x, y)
					if ok && isSymbol(r) {
						// This value counts, let's add it to our total and move on
						completedValue = true
						result += val
						break
					}
				}
				if completedValue {
					break
				}
			}
		}
	}
	return result
}

func isSymbol(r rune) bool {
	// For the purposes of this challenge, a symbol is
	// a non-digit, and not a period.
	return r != '.' && !unicode.IsDigit(r)
}

func findExtent(s string, pos int) (int, int, int) {
	parsingNumbers := false
	start := 0

	if pos > len(s) {
		// We've already reached the end of the string
		return -1, 0, 0
	}

	for i, c := range s[pos:] {
		if unicode.IsDigit(c) {
			if !parsingNumbers {
				parsingNumbers = true
				start = pos + i
			}
		} else if parsingNumbers {
			// We were parsing numbers, and now we're not
			strVal := s[start : pos+i]
			val, _ := strconv.Atoi(strVal)
			return val, start, i + pos - start
		}
	}

	if parsingNumbers {
		// We got to the end of the string, while parsing numbers
		strVal := s[start:]
		val, _ := strconv.Atoi(strVal)
		return val, start, len(s) - start
	}
	return -1, 0, 0
}
