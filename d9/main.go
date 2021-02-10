package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const preambleLen = 25

// count <= 0 means scan indefinitely
func parseInts(
	scanner *bufio.Scanner,
	count int,
	process func(i, num int) (keepParsing bool)) {
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			panic("Invalid input")
		}
		if !process(i, num) {
			return
		}
		i++
		if count > 0 && i >= count {
			break
		}
	}
}

// check that num is a sum of any two numbers in preamble
func verifySum(num int, preamble []int) bool {
	for i := 0; i < len(preamble)-1; i++ {
		for j := i + 1; j < len(preamble); j++ {
			if preamble[i]+preamble[j] == num {
				return true
			}
		}
	}
	return false
}

func part1(numbers []int) int {
	for i := preambleLen; i < len(numbers); i++ {
		if !verifySum(numbers[i], numbers[i-preambleLen:i]) {
			return numbers[i]
		}
	}
	panic("No solution found")
}

func part2(numbers []int) int {
	part1Answer := part1(numbers)
	sum := numbers[0] + numbers[1]
	var i, j int
loop:
	for i = 0; i < len(numbers)-1; {
		for j = 1; j < len(numbers); {
			if sum == part1Answer {
				break loop
			} else if sum < part1Answer {
				j++
				sum += numbers[j]
			} else /* sum > part1Answer */ {
				sum -= numbers[i]
				i++
			}
		}
	}
	if i == len(numbers)-1 && j == len(numbers) {
		panic("No solution found")
	}
	min := numbers[i]
	max := numbers[i]
	for k := i; k <= j; k++ {
		if numbers[k] < min {
			min = numbers[k]
		}
		if numbers[k] > max {
			max = numbers[k]
		}
	}
	return min + max
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
	var answer int

	var numbers []int
	parseInts(bufio.NewScanner(os.Stdin), 0, func(i, num int) bool {
		numbers = append(numbers, num)
		return true
	})

	switch parseVersionArg() {
	case 1:
		answer = part1(numbers)
	case 2:
		answer = part2(numbers)
	}
	fmt.Println(answer)
}
