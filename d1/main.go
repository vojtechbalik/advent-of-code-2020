package main

import (
	"fmt"
	"io"
)

func main() {
	var ints []int

	for {
		var tmp int
		_, err := fmt.Scanf(`%d`, &tmp)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		ints = append(ints, tmp)
	}

	for i := 0; i < len(ints)-2; i++ {
		for j := i + 1; j < len(ints)-1; j++ {
			for k := j + 1; k < len(ints); k++ {
				if ints[i]+ints[j]+ints[k] == 2020 {
					fmt.Printf("%d\n", ints[i]*ints[j]*ints[k])
					break
				}
			}
		}
	}
}
