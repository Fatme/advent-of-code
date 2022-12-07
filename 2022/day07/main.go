package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxSize = 100000

type Directory struct {
	name   string
	dirs   map[string]*Directory
	files  map[string]int
	parent *Directory
}

var root = &Directory{name: "/", dirs: map[string]*Directory{}, files: map[string]int{}, parent: nil}

var mapDirToSize map[string]int = map[string]int{}

func (d *Directory) sum() int {
	sum := 0

	for _, size := range d.files {
		sum += size
	}

	for _, dir := range d.dirs {
		sum += dir.sum()
	}

	mapDirToSize[d.name] = sum

	return sum
}

func isCommand(line string) bool {
	return string(line[0]) == "$"
}

func main() {
	input, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(input), "\n")

	currentDirectory := root

	for _, line := range lines {
		command := strings.Split(line, " ")

		if isCommand(line) {
			commandName := command[1]

			if commandName == "cd" {
				dirName := command[2]

				if dirName == "/" {
					currentDirectory = root
				} else if dirName == ".." {
					currentDirectory = currentDirectory.parent
				} else {
					currentDirectory = currentDirectory.dirs[dirName]
				}
			}
		}

		if command[0] == "dir" {
			dirName := command[1]
			currentDirectory.dirs[dirName] = &Directory{
				dirs:   map[string]*Directory{},
				files:  map[string]int{},
				parent: currentDirectory,
				name:   currentDirectory.name + dirName + "/",
			}
		}

		// isFile
		if fileSize, err := strconv.Atoi(command[0]); err == nil {
			currentDirectory.files[command[1]] = fileSize
		}
	}

	root.sum()

	result := 0
	for _, size := range mapDirToSize {
		if size < maxSize {
			result += size
		}
	}

	fmt.Println(result)
}
