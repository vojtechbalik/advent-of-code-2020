package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	nop = iota
	acc
	jmp
)

type operation struct {
	opType   int
	param    int
	executed bool
}

func interpret(ops []operation) (int, bool) {
	answer := 0
	finished := true
	for i := 0; i < len(ops); i++ {
		if ops[i].executed {
			finished = false
			break
		}
		ops[i].executed = true
		switch ops[i].opType {
		case nop:
			continue
		case acc:
			answer += ops[i].param
		case jmp:
			i += ops[i].param - 1
		default:
			panic("Should be unreachable")
		}
	}
	return answer, finished
}

func (o *operation) swapNopJmp() bool {
	switch o.opType {
	case nop:
		o.opType = jmp
	case jmp:
		o.opType = nop
	default:
		return false
	}
	return true
}

func fixAndInterpret(ops []operation) int {
	for i := 0; i < len(ops); i++ {
		for i := range ops {
			ops[i].executed = false
		}
		if ops[i].swapNopJmp() {
			acc, finished := interpret(ops)
			if finished {
				return acc
			}
			ops[i].swapNopJmp()
		}
	}
	panic("No solution found")
}

func parse(scanner *bufio.Scanner) (ops []operation) {
	regex := regexp.MustCompile(`^(nop|acc|jmp) (\+|-)(\d+)$`)
	for scanner.Scan() {
		line := scanner.Text()
		submatches := regex.FindStringSubmatch(line)
		if submatches == nil {
			panic("invalid format")
		}
		var opType int
		switch submatches[1] {
		case "nop":
			opType = nop
		case "acc":
			opType = acc
		case "jmp":
			opType = jmp
		}
		param, _ := strconv.Atoi(submatches[3])
		if submatches[2] == "-" {
			param *= -1
		}
		ops = append(ops, operation{
			opType:   opType,
			param:    param,
			executed: false,
		})
	}
	return
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
	ops := parse(bufio.NewScanner(os.Stdin))
	var answer int
	switch parseVersionArg() {
	case 1:
		answer, _ = interpret(ops)
	case 2:
		answer = fixAndInterpret(ops)
	}
	fmt.Println(answer)
}
