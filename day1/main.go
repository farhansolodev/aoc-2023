package main

import (
	"fmt"
	"os"
	"time"
)

func timer() func() {
	start := time.Now()
	return func() {
		fmt.Printf("took %v\n", time.Since(start))
	}
}

func main() {
	defer timer()()
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println(part1(file))
	fmt.Println(part2(file))
}
