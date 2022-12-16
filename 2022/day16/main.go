package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseLine(str string) *Node {
	label := str[6:8]

	rateParts := strings.Split(str, " has flow rate=")
	rateRaw := string(rateParts[1])
	rateResult := ""
	index := 0
	for rateRaw[index] != ';' {
		rateResult += string(rateRaw[index])
		index++
	}
	rate, _ := strconv.Atoi(rateResult)

	childrenParts := strings.Split(str, "tunnels lead to valve")
	children := []*Node{}

	if len(childrenParts) > 1 {
		childrenRaw := childrenParts[1]
		if childrenRaw[0] == 's' {
			childrenRaw = childrenRaw[1:]
		}
		childrenRaw = childrenRaw[1:]
		childrenStr := strings.Split(childrenRaw, ", ")
		for _, child := range childrenStr {
			children = append(children, &Node{label: child})
		}
	} else {
		parts := strings.Split(str, "tunnel leads to valve ")
		children = append(children, &Node{label: parts[1]})
	}

	return &Node{label: label, rate: rate, children: children}
}

type Node struct {
	label    string
	children []*Node
	rate     int
}

var nodes = map[string]*Node{}

type State struct {
	node             *Node
	openedNodes      map[string]bool
	minutes          int
	releasedPressure int
}

func (s *State) id() string {
	return fmt.Sprintf("%s%d", s.node.label, s.releasedPressure)
}

var states = map[string]bool{}

func createState(node *Node, openedNodes map[string]bool, minutes int, releasedPresure int) *State {
	state := &State{node: node, openedNodes: openedNodes, minutes: minutes, releasedPressure: releasedPresure}

	if states[state.id()] {
		return nil
	}

	states[state.id()] = true

	return state
}

func main() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		node := parseLine(line)
		nodes[node.label] = node
	}

	initialNode := nodes["AA"]
	initialState := createState(initialNode, map[string]bool{}, 30, 0)

	max := 0

	queue := []*State{}
	queue = append(queue, initialState)

	addToQueue := func(state *State) {
		if state != nil {
			queue = append(queue, state)
		}
	}

	for len(queue) > 0 {
		currentState := queue[0]
		currentNode := currentState.node

		queue = queue[1:]

		max = int(math.Max(float64(max), float64(currentState.releasedPressure)))

		if !currentState.openedNodes[currentNode.label] && currentNode.rate > 0 {
			minutes := currentState.minutes - 1
			releasedPressure := currentState.releasedPressure + (currentNode.rate * minutes)

			openedNodes := map[string]bool{}
			for key, value := range currentState.openedNodes {
				openedNodes[key] = value
			}
			openedNodes[currentNode.label] = true

			newState := createState(currentNode, openedNodes, minutes, releasedPressure)
			addToQueue(newState)
		}

		for _, child := range currentNode.children {
			newState := createState(nodes[child.label], currentState.openedNodes, currentState.minutes-1, currentState.releasedPressure)
			addToQueue(newState)
		}
	}

	fmt.Println(max)
}
