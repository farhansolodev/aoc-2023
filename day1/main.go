package main

import (
	"fmt"
	"os"
	"time"
)

func printer(name string) func(any) {
	start := time.Now()
	return func(v any) {
		fmt.Printf("%s: %v\n", name, v)
		fmt.Printf("-- took %v\n", time.Since(start))
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	printPart1 := printer("part1")
	p1sum := part1(file)
	printPart1(p1sum)

	printPart2 := printer("part2")
	p2sum := part2(file)
	printPart2(p2sum)
}
