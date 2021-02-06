package main

import (
	"fmt"
	"io"
)

func main() {
	var pwd string
	var c byte
	var first, second int
	scanFormat := "%d-%d %c: %s"
	validPwdCount := 0

	_, err := fmt.Scanf(scanFormat, &first, &second, &c, &pwd)
	for err == nil {
		if (pwd[first-1] == c) != (pwd[second-1] == c) {
			validPwdCount++
		}
		_, err = fmt.Scanf(scanFormat, &first, &second, &c, &pwd)
	}
	if err != io.EOF {
		panic(err)
	}
	fmt.Printf("%d\n", validPwdCount)
}
