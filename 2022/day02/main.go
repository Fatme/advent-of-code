package main

import (
	"fmt"
	"strings"
)

const Rock = 1
const Paper = 2
const Scissors = 3

var data = map[string]int{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win
func main() {
	sum := 0

	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i++ {
		parts := strings.Split(lines[i], " ")

		player := data[strings.ToUpper(parts[0])]
		result := parts[1]

		if result == "Y" {
			sum += player + 3
		}

		if result == "X" {
			// lose
			if player == 1 {
				sum += 3
			}
			if player == 2 {
				sum += 1
			}
			if player == 3 {
				sum += 2
			}
		}

		if result == "Z" {
			// win
			sum += 6
			if player == 1 {
				sum += 2
			}
			if player == 2 {
				sum += 3
			}
			if player == 3 {
				sum += 1
			}
		}
	}

	fmt.Println(sum)
}

// Part1
func main1() {
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
