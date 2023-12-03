package main

import (
	"fmt"
	"os"
	"time"
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
	const filepath = "big.txt"
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	printPart1 := printer("part1")
	p1sum := part1(file)
	printPart1(p1sum)
	file.Close()

	file, err = os.Open(filepath)
	if err != nil {
		panic(err)
	}
	printPart2 := printer("part2")
	p2sum := part2(file)
	printPart2(p2sum)
	file.Close()
}
