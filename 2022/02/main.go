// I tried to be "smart" for this one and based the solver on subtracting the
// ASCII value of the players play from the opponents play, which results in one
// of 5 values between -2 and 2, which can than be mapped to the game outcome
//
// As for the second part I was pretty lazy and just added a LUT that maps the
// instructions to the needed play
package main

import (
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("in.txt")

	if err != nil {
		println(err)
		return
	}

	println("Part 1: ", part1(data))
	println("Part 2: ", part2(data))
}

func part1(data []byte) int {
	lines := strings.Split(string(data), "\n")

	score := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		moves := strings.Split(line, " ")

		a := int(moves[0][0])
		b := int(moves[1][0]) - 23

		score += solveGame(a, b)
	}

	return score
}

// LUT mapping the instruction to the play needed to make it happen
var LUT = map[int]map[int]int {
	'A': {
		'X': 'C',
		'Y': 'A',
		'Z': 'B',
	},
	'B': {
		'X': 'A',
		'Y': 'B',
		'Z': 'C',
	},
	'C': {
		'X': 'B',
		'Y': 'C',
		'Z': 'A',
	},
}

func part2(data []byte) int {
	lines := strings.Split(string(data), "\n")

	score := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		instructions := strings.Split(line, " ")

		a := int(instructions[0][0])
		b := int(instructions[1][0])

		// Look up what "shape" we have to play from the LUT
		toPlay := LUT[a][b]

		score += solveGame(a, toPlay)
	}

	return score
}

// Resolves the scores for a single game, the inputs are assumed to be the ASCII
// codes for A, B or C, meaning Rock, Paper and Scissors
func solveGame(a int, b int) (score int) {

	// Add the "shape" score
	score += b - 64

	switch a - b {

	// Win
	case -1, 2:
		score += 6
		break

	// Draw
	case 0:
		score += 3
		break

	// Loss
	case -2, 1:
		break
	}

	return score
}
