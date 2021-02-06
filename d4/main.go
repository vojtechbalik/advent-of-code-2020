package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func validateRecord1(r map[string]string) bool {
	required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	var i int
	for i = 0; i < len(required); i++ {
		if _, present := r[required[i]]; !present {
			break
		}
	}
	return i == len(required)
}

func validateRecord2(r map[string]string) bool {
	if !validateRecord1(r) {
		return false
	}

	for k, v := range r {
		switch k {
		case "byr":
			if byr, err := strconv.ParseInt(v, 10, 0); err != nil || !(1920 <= byr && byr <= 2002) {
				return false
			}
		case "iyr":
			if byr, err := strconv.ParseInt(v, 10, 0); err != nil || !(2010 <= byr && byr <= 2020) {
				return false
			}
		case "eyr":
			if byr, err := strconv.ParseInt(v, 10, 0); err != nil || !(2020 <= byr && byr <= 2030) {
				return false
			}
		case "hgt":
			var num int
			var units string
			fmt.Sscanf(v, "%d%s", &num, &units)
			switch units {
			case "cm":
				if !(150 <= num && num <= 193) {
					return false
				}
			case "in":
				if !(59 <= num && num <= 76) {
					return false
				}
			default:
				return false
			}
		case "hcl":
			validHcl := regexp.MustCompile(`^#[0-9a-f]{6}$`)
			if !validHcl.MatchString(v) {
				return false
			}
		case "ecl":
			validEcl := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
			if !validEcl.MatchString(v) {
				return false
			}
		case "pid":
			validPid := regexp.MustCompile(`^[0-9]{9}$`)
			if !validPid.MatchString(v) {
				return false
			}
		}
	}
	return true
}

func main() {
	var validateRecord func(map[string]string) bool
	if len(os.Args) < 2 {
		panic("Missing argument")
	}
	version, err := strconv.ParseInt(os.Args[1], 10, 0)
	if err != nil {
		panic("Invalid argument")
	}
	switch version {
	case 1:
		validateRecord = validateRecord1
	case 2:
		validateRecord = validateRecord2
	default:
		panic("Invalid argument")
	}

	scanner := bufio.NewScanner(os.Stdin)
	rec := make(map[string]string)
	validRecordsCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if validateRecord(rec) {
				validRecordsCount++
			}
			rec = make(map[string]string)
		}
		items := strings.Fields(line)
		for _, item := range items {
			split := strings.SplitN(item, ":", 2)
			rec[split[0]] = split[1]
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	// there might not be a blank line after last record
	if validateRecord(rec) {
		validRecordsCount++
	}
	fmt.Printf("%d\n", validRecordsCount)
}
