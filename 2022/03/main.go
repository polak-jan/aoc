// This was a fun one. The way I solved it is by converting the input into
// bitfields and than finding the common items using bitwise AND. Not much more
// to it surprisingly.
package main

import (
	"math"
	"os"
	"strings"
)

func main() {
	println("AoC 2022 Day 3")
	println("https://adventofcode.com/2022/day/3")

	data, err := os.ReadFile("in.txt")

	if err != nil {
		println(err)
		return
	}

	println("Part 1: ", part1(data))
	println("Part 2: ", part2(data))
}

// Convert the input string into a bitfield of the items present in it. The
// values are normalised into the priority specified by the instructions
func toBitField(in string) uint64 {

	var field uint64 = 0

	for _, char := range in {

		// Move the uppercase letters after the lower case ones
		if char < 95 {
			char += 32 + 26
		}

		// Move the entire scale to start at 0
		char -= 96

		field = field | 1 << char
	}

	return field
}

func part1(data []byte) int {
	lines := strings.Split(string(data), "\n")

	total := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		// Divide the line into two halves and bitwise AND them to find
		// the common items between them
		common :=
			toBitField(line[:len(line)/2]) &
			toBitField(line[len(line)/2:])

		// Find out which bits are set on the bitfield and sum them up
		var i uint64
		for i = 0; i < 53; i++ {
			if (1 << i) & common > 0 {
				total += int(i)
			}
		}
	}

	return total
}

func part2(data []byte) int {
	lines := strings.Split(string(data), "\n")

	total := 0

	// Holds the common items between each group of 3, starts as a fully
	// filled bitfield
	var aggregate uint64 = math.MaxUint64

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}

		aggregate = aggregate & toBitField(line)

		// Find the item common to the previous three lines and add it
		// to the total
		if (i + 1) % 3 == 0 {
			var j uint64

			for j = 0; j < 53; j++ {
				if (1 << j) & aggregate > 0 {
					total += int(j)
				}
			}

			aggregate = math.MaxUint64
		}
	}

	return total
}
