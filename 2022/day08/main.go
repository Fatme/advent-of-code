package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func process(board [][]int, i, j int) (bool, int) {
	isVisibleFromLeft, leftCount := processLeft(board, i, j)
	isVisibleFromRight, rightCount := processRight(board, i, j)
	isVisibleFromTop, topCount := processTop(board, i, j)
	isVisibleFromBottom, bottomCount := processBottom(board, i, j)

	isVisible := isVisibleFromLeft || isVisibleFromRight || isVisibleFromTop || isVisibleFromBottom
	score := leftCount * rightCount * topCount * bottomCount

	return isVisible, score
}

func processLeft(board [][]int, i, j int) (bool, int) {
	visible := true
	count := 0

	for k := j - 1; k >= 0; k-- {
		count++
		if board[i][k] >= board[i][j] {
			visible = false
			break
		}
	}

	return visible, count
}

func processRight(board [][]int, i, j int) (bool, int) {
	visible := true
	count := 0

	for k := j + 1; k < len(board[i]); k++ {
		count++
		if board[i][k] >= board[i][j] {
			visible = false
			break
		}
	}

	return visible, count
}

func processTop(board [][]int, i, j int) (bool, int) {
	visible := true
	count := 0

	for k := i - 1; k >= 0; k-- {
		count++
		if board[k][j] >= board[i][j] {
			visible = false
			break
		}
	}

	return visible, count
}

func processBottom(board [][]int, i, j int) (bool, int) {
	visible := true
	count := 0

	for k := i + 1; k < len(board); k++ {
		count++
		if board[k][j] >= board[i][j] {
			visible = false
			break
		}
	}

	return visible, count
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

	maxScore := 0

	for i := 1; i < len(board)-1; i++ {
		for j := 1; j < len(board[i])-1; j++ {
			isVisible, score := process(board, i, j)
			if isVisible {
				count++
			}

			if score > maxScore {
				maxScore = score
			}
		}
	}

	fmt.Println(count)
	fmt.Println(maxScore)
}
