package day5

import (
	"math"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/yeurch/aoc2023/util"
)

func parseAlmanac(almanacLines []string) [][][]int {
	almanac := make([][][]int, 0)
	var mapping [][]int
	for _, line := range almanacLines {
		if len(line) == 0 {
			continue // skip blank lines
		}
		firstChar, _ := utf8.DecodeRuneInString(line)
		if !unicode.IsNumber(firstChar) {
			if len(mapping) > 0 {
				almanac = append(almanac, mapping)
			}
			mapping = make([][]int, 0)
			continue
		}

		ruleValues := strings.Fields(line)
		v0, _ := strconv.Atoi(ruleValues[0])
		v1, _ := strconv.Atoi(ruleValues[1])
		v2, _ := strconv.Atoi(ruleValues[2])
		rule := []int{v0, v1, v2}
		mapping = append(mapping, rule)
	}
	almanac = append(almanac, mapping) // append our last mapping

	return almanac
}

func doImpl(seedNums util.Set[int], almanacLines []string) int {
	almanac := parseAlmanac(almanacLines)

	result := math.MaxInt
	for val := range seedNums.Iter() {
		for _, mapping := range almanac {
			for _, rule := range mapping {
				if val >= rule[1] && val < rule[1]+rule[2] {
					val = rule[0] + val - rule[1]
					break
				}
			}
		}
		if val < result {
			result = val
		}
	}

	return result
}

func lookup(val int, almanac [][][]int) int {
	for _, mapping := range almanac {
		for _, rule := range mapping {
			if val >= rule[1] && val < rule[1]+rule[2] {
				val = rule[0] + val - rule[1]
				break
			}
		}
	}
	return val
}

func reverseLookup(val int, mapping [][]int) int {
	for _, rule := range mapping {
		if val >= rule[1] && val < rule[1]+rule[2] {
			return rule[1] + val - rule[0]
		}
	}
	return val
}
