package main

import (
	"fmt"
	"strings"
)

const Rock = 1
const Paper = 2
const Scissors = 3

func main() {
	data := map[string]int{
		"A": Rock,
		"B": Paper,
		"C": Scissors,
		"X": Rock,
		"Y": Paper,
		"Z": Scissors,
	}

	sum := 0

	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i++ {
		parts := strings.Split(lines[i], " ")

		player1 := data[strings.ToUpper(parts[0])]
		player2 := data[strings.ToUpper(parts[1])]

		sum += player2

		if player1 == player2 {
			sum += 3
		}

		isWinForPlayer2 := (player2 == Rock && player1 == Scissors) ||
			(player2 == Scissors && player1 == Paper) ||
			(player2 == Paper && player1 == Rock)

		if isWinForPlayer2 {
			sum += 6
		}
	}

	fmt.Println(sum)
}
