package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isVisible(board [][]int, i, j int) bool {
	return isVisibleFromLeft(board, i, j) || isVisibleFromRight(board, i, j) || isVisibleFromTop(board, i, j) || isVisibleFromBottom(board, i, j)
}

func isVisibleFromLeft(board [][]int, i, j int) bool {
	visible := true

	for k := j - 1; k >= 0; k-- {
		if board[i][k] >= board[i][j] {
			visible = false
			break
		}
	}

	return visible
}

func isVisibleFromRight(board [][]int, i, j int) bool {
	visible := true

	for k := j + 1; k < len(board[i]); k++ {
		if board[i][k] >= board[i][j] {
			visible = false
			break
		}
	}

	return visible
}

func isVisibleFromTop(board [][]int, i, j int) bool {
	visible := true

	for k := i - 1; k >= 0; k-- {
		if board[k][j] >= board[i][j] {
			visible = false
			break
		}
	}

	return visible
}

func isVisibleFromBottom(board [][]int, i, j int) bool {
	visible := true

	for k := i + 1; k < len(board); k++ {
		if board[k][j] >= board[i][j] {
			visible = false
			break
		}
	}

	return visible
}

func main() {
	input, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(input), "\n")

	board := make([][]int, len(lines))

	for index, line := range lines {
		board[index] = make([]int, len(line))
		for j := 0; j < len(line); j++ {
			num, _ := strconv.Atoi(string(line[j]))
			board[index][j] = num
		}
	}

	count := 2*(len(board)+len(board[0])) - 4

	for i := 1; i < len(board)-1; i++ {
		for j := 1; j < len(board[i])-1; j++ {
			if isVisible(board, i, j) {
				count++
			}
		}
	}

	fmt.Println(count)
}
