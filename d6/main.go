package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strconv"
)

func combineAnswers1(a1, a2 uint32) uint32 {
	return a1 | a2
}

func initAnswers1() uint32 {
	return 0
}

func combineAnswers2(a1, a2 uint32) uint32 {
	return a1 & a2
}

func initAnswers2() uint32 {
	return 1<<32 - 1
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
	var initAnswers func() uint32
	var combineAnswers func(uint32, uint32) uint32
	switch version {
	case 1:
		initAnswers = initAnswers1
		combineAnswers = combineAnswers1
	case 2:
		initAnswers = initAnswers2
		combineAnswers = combineAnswers2
	default:
		panic("Unreachable")
	}

	answersSum := 0
	var answers uint32 = initAnswers()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var currentAnswer uint32 = 0
		if line == "" {
			answersSum += bits.OnesCount32(answers)
			answers = initAnswers()
			continue
		}
		for _, q := range line {
			currentAnswer |= 1 << (q - 'a')
		}
		answers = combineAnswers(answers, currentAnswer)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	answersSum += bits.OnesCount32(answers)

	fmt.Printf("%d\n", answersSum)
}
