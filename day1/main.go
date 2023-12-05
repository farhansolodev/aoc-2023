package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	day1 "github.com/farhansolodev/aoc-2023/src"
)

func timer() func() time.Duration {
	start := time.Now()
	return func() time.Duration {
		return time.Since(start)
	}
}

func main() {
	filepath := flag.String("file", "input.txt", "path to input file")
	flag.Parse()

	// part 1
	file, err := os.Open(*filepath)
	if err != nil {
		panic(err)
	}
	getPart1Duration := timer()
	p1sum := day1.Part1(file)
	p1d := getPart1Duration()
	file.Close()
	fmt.Printf("%s: %v\n", "part1", p1sum)
	fmt.Printf("-- took %v\n", p1d)

	// part 2
	file, err = os.Open(*filepath)
	if err != nil {
		panic(err)
	}
	getPart2Duration := timer()
	p2sum := day1.Part2(file)
	p2d := getPart2Duration()
	file.Close()
	fmt.Printf("%s: %v\n", "part2", p2sum)
	fmt.Printf("-- took %v\n", p2d)

	fmt.Printf("\nPart 2 took %fx longer", float32(p2d.Microseconds())/float32(p1d.Microseconds()))
}
