package day3

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yeurch/aoc2023/challenge"
)

func bCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "b",
		Short: "Day 3, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", partB(challenge.FromFile()))
		},
	}
}

type gearMap map[struct {
	x int
	y int
}][]int

func partB(challenge *challenge.Input) int {
	gears := make(gearMap, 0)
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

			for y := ln - 1; y < ln+2; y++ {
				for x := start - 1; x < start+length+1; x++ {
					r, ok := grid.TileAt(x, y)
					if ok && r == '*' {
						// This value counts, let's add it to our total and move on
						gearPos := struct {
							x int
							y int
						}{x, y}
						if _, ok := gears[gearPos]; !ok {
							gears[gearPos] = make([]int, 0)
						}
						gears[gearPos] = append(gears[gearPos], val)
					}
				}
			}
		}
	}

	result := 0
	for gear := range gears {
		if len(gears[gear]) == 2 {
			result += gears[gear][0] * gears[gear][1]
		}
	}
	return result
}
