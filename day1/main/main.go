package main

import (
	"fmt"
	"os"
	"time"

	"github.com/farhansolodev/aoc-2023/day1"
)

func printer(name string) func(any) {
	start := time.Now()
	return func(v any) {
		duration := time.Since(start)
		fmt.Printf("%s: %v\n", name, v)
		fmt.Printf("-- took %v\n", duration)
	}
}

func main() {
	const filepath = "input.txt"

	// part 1
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	printPart1 := printer("part1")
	p1sum := day1.Part1(file)
	printPart1(p1sum)
	file.Close()

	// part 2
	file, err = os.Open(filepath)
	if err != nil {
		panic(err)
	}
	printPart2 := printer("part2")
	p2sum := day1.Part2(file)
	printPart2(p2sum)
	file.Close()
}
