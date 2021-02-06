package main

import (
	"fmt"
	"io"
)

func countTrees(right, down int, trees []string) (n int) {
	for i := 0; i < len(trees); i += down {
		if trees[i][(i*right/down)%len(trees[i])] == '#' {
			n++
		}
	}
	return
}

func main() {
	var trees []string
	scanFormat := "%s"

	var treeLine string
	_, err := fmt.Scanf(scanFormat, &treeLine)
	for err == nil {
		trees = append(trees, treeLine)
		_, err = fmt.Scanf(scanFormat, &treeLine)
	}
	if err != io.EOF {
		panic(err)
	}

	slopes := []struct {
		right, down int
	}{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	product := 1
	for _, slope := range slopes {
		product *= countTrees(slope.right, slope.down, trees)
	}
	fmt.Printf("%d\n", product)
}
