package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"../queue"
)

func parse(scanner *bufio.Scanner, processLine func(bag string, contents []bagItem)) {
	regexBag := regexp.MustCompile(`^(\w+ \w+) bags contain (no other bags|[^\.]*)\.$`)
	regexContained := regexp.MustCompile(`^\s*(\d+) (\w+ \w+) bag(s?)\s*$`)
	for scanner.Scan() {
		line := scanner.Text()
		submatches := regexBag.FindStringSubmatch(line)
		bag := submatches[1]
		if submatches[2] == "no other bags" {
			processLine(bag, []bagItem{})
			continue
		}
		allBagItemsSplit := strings.Split(submatches[2], ",")
		var bagItems []bagItem
		for _, split := range allBagItemsSplit {
			submatches := regexContained.FindStringSubmatch(split)
			qty, err := strconv.Atoi(submatches[1])
			if err != nil {
				panic(err)
			}
			bagItems = append(bagItems, bagItem{
				qty:   qty,
				color: submatches[2]})
		}
		processLine(bag, bagItems)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func insertIntoGraph(graph map[string]map[string]struct{}, from, to string) {
	neighbors, ok := graph[from]
	if !ok {
		neighbors = make(map[string]struct{})
		graph[from] = neighbors
	}
	neighbors[to] = struct{}{}
}

func findReachable(
	graph map[string]map[string]struct{},
	start string) (r map[string]struct{}) {
	// BFS
	q := queue.New()
	r = make(map[string]struct{})
	q.Push(start)
	for q.Len() > 0 {
		node, _ := q.Pop()
		nodeS := node.(string)
		r[nodeS] = struct{}{}
		for neighbor := range graph[nodeS] {
			if _, isClosed := r[neighbor]; !isClosed {
				q.Push(neighbor)
			}
		}
	}
	return
}

// I treat the problem as a directed graph, where nodes are bag colors.

// For part 1, edge from node1 to node2 means 'node1 can be contained by
// node2'. Solution is then the cardinality of nodes, that are reachable from
// from 'shiny gold'.
func part1(scanner *bufio.Scanner) int {
	graph := make(map[string]map[string]struct{})
	parse(scanner, func(bag string, contents []bagItem) {
		for _, item := range contents {
			insertIntoGraph(graph, item.color, bag)
		}
	})
	return len(findReachable(graph, "shiny gold")) - 1
}

type bagItem struct {
	qty   int
	color string
}

// For part 2, edges are reversed and also have associated quantity.
// The description also implies the graph has no cycles (otherwise the solution
// could be infinite).
func part2(scanner *bufio.Scanner) int {
	graph := make(map[string][]bagItem)
	parse(scanner, func(bag string, contents []bagItem) {
		graph[bag] = contents
	})

	bagCount := 0
	q := queue.New()
	for _, neighbor := range graph["shiny gold"] {
		q.Push(neighbor)
	}
	for q.Len() > 0 {
		node, _ := q.Pop()
		nodeT := node.(bagItem)
		bagCount += nodeT.qty
		for _, neighbor := range graph[nodeT.color] {
			q.Push(bagItem{
				qty:   nodeT.qty * neighbor.qty,
				color: neighbor.color})
		}
	}
	return bagCount
}

func parseVersionArg() int {
	if len(os.Args) < 2 {
		panic("Missing argument")
	}
	version, err := strconv.Atoi(os.Args[1])
	if err != nil || !(1 <= version && version <= 2) {
		panic("Invalid argument")
	}
	return version
}

func main() {
	version := parseVersionArg()
	scanner := bufio.NewScanner(os.Stdin)
	var answer int
	switch version {
	case 1:
		answer = part1(scanner)
	case 2:
		answer = part2(scanner)
	default:
		panic("Should be unreachable")
	}
	fmt.Println(answer)
}
