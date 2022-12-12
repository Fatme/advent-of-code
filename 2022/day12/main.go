package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/albertorestifo/dijkstra"
)

type Graph struct {
	grid [][]*Pair
}

func (g *Graph) getValue(ch byte) int {
	value := int(ch)

	if ch == 'S' {
		value = 96
	} else if ch == 'E' {
		value = 123
	}

	return value
}

func (g *Graph) populate(lines []string) (dijkstra.Graph, *Pair, *Pair, []*Pair) {
	graph := dijkstra.Graph{}

	g.grid = make([][]*Pair, len(lines))

	startPair := &Pair{}
	endPair := &Pair{}

	var starts []*Pair

	for index, line := range lines {
		g.grid[index] = make([]*Pair, len(line))
		for j := 0; j < len(line); j++ {
			value := g.getValue(line[j])

			pair := &Pair{i: index, j: j, value: value}

			if line[j] == 'S' {
				startPair = pair
			} else if line[j] == 'E' {
				endPair = pair
			} else if line[j] == 'a' {
				starts = append(starts, pair)
			}

			g.grid[index][j] = pair
		}
	}

	for i := range g.grid {
		for j := range g.grid[i] {
			current := g.grid[i][j]
			if _, found := graph[current.String()]; !found {
				graph[current.String()] = g.getNextMoves(current)
			}
		}
	}

	return graph, startPair, endPair, starts
}

func (g *Graph) isInGrid(i, j int) bool {
	return i >= 0 && i < len(g.grid) && j >= 0 && j < len(g.grid[0])
}

func (g *Graph) getNextMoves(pair *Pair) map[string]int {
	moves := map[string]int{}

	process := func(i, j int) {
		if g.isInGrid(i, j) {
			newPair := g.grid[i][j]
			if pair.isAllowedMove(newPair) {
				moves[newPair.String()] = 1
			}
		}
	}

	process(pair.i-1, pair.j)
	process(pair.i+1, pair.j)
	process(pair.i, pair.j-1)
	process(pair.i, pair.j+1)

	return moves
}

func (p *Pair) isAllowedMove(new *Pair) bool {
	return (new.value - p.value) <= 1
}

type Pair struct {
	i     int
	j     int
	value int
}

func (p *Pair) String() string {
	return fmt.Sprintf("%d,%d", p.i, p.j)
}

func main() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")

	graph, startPair, endPair, starts := (&Graph{}).populate(lines)

	_, part1, _ := graph.Path(startPair.String(), endPair.String())

	part2 := math.MaxInt
	for _, s := range starts {
		_, cost, _ := graph.Path(s.String(), endPair.String())
		if cost < part2 && cost > 0 {
			part2 = cost
		}
	}

	fmt.Println(part1)
	fmt.Println(part2)
}
