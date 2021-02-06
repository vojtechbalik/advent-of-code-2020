package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

func conversionFunc(zero, one rune) func(rune) rune {
	return func(r rune) rune {
		switch r {
		case zero:
			return '0'
		case one:
			return '1'
		default:
			panic("Unexpected rune")
		}
	}
}

func main() {
	version := parseVersionArg()

	var IDs []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		seat := scanner.Text()
		row, err := strconv.ParseInt(
			strings.Map(conversionFunc('F', 'B'), seat[:7]), 2, 0)
		if err != nil {
			panic("This should be unreachable")
		}
		col, err := strconv.ParseInt(
			strings.Map(conversionFunc('L', 'R'), seat[7:]), 2, 0)
		if err != nil {
			panic("This should be unreachable")
		}
		IDs = append(IDs, int(row)*8+int(col))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	sort.Ints(IDs)

	if version == 1 {
		fmt.Printf("%d\n", IDs[len(IDs)-1])
		return
	}

	for i := 0; i < len(IDs)-1; i += 2 {
		if suspect := IDs[i] + 1; IDs[i+1] != suspect {
			fmt.Printf("%d\n", suspect)
			return
		}
	}
}
