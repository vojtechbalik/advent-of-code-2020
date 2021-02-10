package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func appendInts(scanner *bufio.Scanner, into []int) []int {
	var ints []int = into
	for scanner.Scan() {
		line := scanner.Text()
		parsed, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalln("Could not parse int", err)
		}
		ints = append(ints, parsed)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln("Error parsing input", err)
	}
	return ints
}

func part1(outputs []int) int {
	oneJoltDiff := 0
	threeJoltDiff := 0

	for i := 0; i < len(outputs)-1; i++ {
		diff := outputs[i+1] - outputs[i]
		switch diff {
		case 1:
			oneJoltDiff++
		case 2:
			continue
		case 3:
			threeJoltDiff++
		default:
			log.Fatalln("Diff not in [1, 3]")
		}
	}
	return oneJoltDiff * threeJoltDiff
}

func part2(outputs []int) int {
	// To hold the number of possible arrangements starting from each index. (The
	// very last one is meaningless, as the one before it is present in
	// every arrangement.)
	arrs := make([]int, len(outputs))
	arrs[len(arrs)-2] = 1
	arrs[len(arrs)-3] = 1
	for i := len(arrs) - 4; i >= 0; i-- {
		for j := 1; j <= 3; j++ {
			if outputs[i+j]-outputs[i] <= 3 {
				arrs[i] += arrs[i+j]
			}
		}
	}
	return arrs[0]
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
	outputs := []int{0}
	outputs = appendInts(bufio.NewScanner(os.Stdin), outputs)
	sort.Ints(outputs)
	outputs = append(outputs, outputs[len(outputs)-1]+3)

	var answer int
	switch parseVersionArg() {
	case 1:
		answer = part1(outputs)
	case 2:
		answer = part2(outputs)
	}
	fmt.Println(answer)

}
