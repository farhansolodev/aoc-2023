package main

import (
	"fmt"
	"time"
)

func timer() func() {
	start := time.Now()
	return func() {
		fmt.Printf("took %v\n", time.Since(start))
	}
}

func main() {
	// part1()
	part2()
}
