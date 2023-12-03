package main

import (
	"fmt"
	"os"
	"time"
)

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	part1Lap := timer("part1")
	fmt.Println(part1(file))
	part1Lap()

	part2Lap := timer("part2")
	fmt.Println(part2(file))
	part2Lap()
}
