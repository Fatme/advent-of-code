package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type item [2]int

const decryptionKey = 811589153

func main() {
	lines := strings.Split(input, "\n")

	solvePart1(lines)
	solvePart2(lines)
}

func readInput(lines []string, multiplicator int) ([]item, []int) {
	items := []item{}
	nums := []int{}

	for index, line := range lines {
		num, _ := strconv.Atoi(line)
		num *= multiplicator
		nums = append(nums, num)
		items = append(items, item{num, index})
	}

	return items, nums
}

func mix(items []item, nums []int, length int) {
	for index, num := range nums {
		if num == 0 {
			continue
		}

		currentPosition := 0
		for i := 0; i < length; i++ {
			if index == items[i][1] && num == items[i][0] {
				currentPosition = i
				break
			}
		}

		newPosition := getNewPosition(currentPosition, num, length)
		if newPosition == 0 {
			newPosition = length - 1
		}

		if newPosition > currentPosition {
			copy(items[currentPosition:], items[currentPosition+1:newPosition+1])
		} else {
			copy(items[newPosition+1:], items[newPosition:currentPosition])
		}

		items[newPosition][0] = num
		items[newPosition][1] = index
	}
}

func solvePart1(lines []string) {
	length := len(lines)
	items, nums := readInput(lines, 1)

	mix(items, nums, length)

	sum := getResult(items, nums, length)
	fmt.Println(sum)
}

func solvePart2(lines []string) {
	length := len(lines)
	items, nums := readInput(lines, decryptionKey)

	for i := 0; i < 10; i++ {
		mix(items, nums, length)
	}

	sum := getResult(items, nums, length)
	fmt.Println(sum)
}

func getResult(items []item, nums []int, length int) int {
	initialZeroIndex := 0

	for index, item := range items {
		if item[0] == 0 {
			initialZeroIndex = index
			break
		}
	}

	num1000 := items[(initialZeroIndex+1000)%length][0]
	num2000 := items[(initialZeroIndex+2000)%length][0]
	num3000 := items[(initialZeroIndex+3000)%length][0]

	sum := num1000 + num2000 + num3000

	return sum
}

func getNewPosition(currentPosition, num, length int) int {
	result := (currentPosition + num) % (length - 1)

	if num > 0 {
		return result
	}

	return (result + (length - 1)) % (length - 1)
}
