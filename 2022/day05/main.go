package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/lists/arraylist"
)

const COLS = 9
const INPUT_FILE = "input.txt"

func parseInput(lines []string) [COLS]*arraylist.List {
	var data [COLS]*arraylist.List

	for _, line := range lines {
		if !isInstructionLine(line) && line != "" && !strings.Contains(line, " 1   2") {
			col := 0

			for j := 0; j < len(line); j += 4 {
				item := string(line[j+1])
				// Ignore empty cells
				if item == " " {
					col++
					continue
				}

				if data[col] == nil {
					data[col] = arraylist.New()
				}

				data[col].Insert(0, item)
				col++
			}
		}
	}

	return data
}

func isInstructionLine(line string) bool {
	return strings.Contains(line, "move")
}

func parseInstructionLine(line string) (int, int, int) {
	parts := strings.Split(line, " ")

	movesCount, _ := strconv.Atoi(parts[1])
	from, _ := strconv.Atoi(parts[3])
	to, _ := strconv.Atoi(parts[5])

	return movesCount, from, to
}

func main() {
	input, _ := os.ReadFile(INPUT_FILE)

	lines := strings.Split(string(input), "\n")

	data := parseInput(lines)

	for _, line := range lines {
		if isInstructionLine(line) {
			movesCount, from, to := parseInstructionLine(line)

			fromList := data[from-1]
			toList := data[to-1]

			for i := movesCount - 1; i >= 0; i-- {
				fromIndex := fromList.Size() - 1
				item, _ := fromList.Get(fromIndex)
				fromList.Remove(fromIndex)
				toList.Add(item)
			}
		}
	}

	result := ""

	for i := 0; i < len(data); i++ {
		item, _ := data[i].Get(data[i].Size() - 1)
		result += item.(string)
	}

	fmt.Println(result)
}
