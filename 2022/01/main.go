// My solution to this one is to simply create a list of all the elves, sort
// them, and then just take the one with the most calories for the first part,
// and sum up the top 3 for the second part. Not much room for creativity here.
package main

import (
	"os"
	"sort"
	"strconv"
	"strings"
)


func main() {
	println("AoC 2022 Day 1")
	println("https://adventofcode.com/2022/day/1")

	data, err := os.ReadFile("in.txt")

	if err != nil {
		println(err)
		return
	}

	lines := strings.Split(string(data), "\n")

	// List of all elves present in the input
	elves := make([]int, 0)

	accumulator := 0

	// Go through all the lines in the file and aggregate all the elves into
	// a single slice
	for _, line := range lines {

		// Empty line is a separator between elves, so we push the
		// current sum into the slice and go back to 0
		if len(line) == 0 {
			elves = append(elves, accumulator)
			accumulator = 0

			continue
		}

		// Parse the line
		parsed, err := strconv.Atoi(line)

		if err != nil {
			println(err)
			return
		}

		// Add the current line to the "accumulator"
		accumulator += parsed
	}

	// Sort the elves, with the most calories being first
	sort.Sort(sort.Reverse(sort.IntSlice(elves)))

	{ // Part 1
		println("Part 1: ", elves[0])
	}

	{ // Part 2
		total := 0

		for _, elf := range elves[:3] {
			total += elf
		}

		println("Part 2: ", total)
	}
}
