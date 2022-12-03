package main

import "os"

func main() {
	println("AoC 2022 Day 2")
	println("https://adventofcode.com/2022/day/2")

	data, err := os.ReadFile("in.txt")

	if err != nil {
		println(err)
		return
	}

	println("Part 1: ", part1(data))
	println("Part 2: ", part2(data))
}

func part1(data []byte) int {
	return 0
}

func part2(data []byte) int {
	return 0
}
