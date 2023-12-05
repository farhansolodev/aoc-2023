package main

import (
	"os"
	"testing"

	day1 "github.com/farhansolodev/aoc-2023/src"
)

var result int

func benchmark(input string, part int, b *testing.B) {
	file, err := os.Open("../" + input)
	if err != nil {
		panic(err)
	}
	var r int
	switch part {
	case 1:
		for i := 0; i < b.N; i++ {
			r = day1.Part1(file)
		}
	case 2:
		for i := 0; i < b.N; i++ {
			r = day1.Part2(file)
		}
	default:
		panic("invalid part provided")
	}
	result = r
}

func BenchmarkPart1_inputTXT(b *testing.B) { benchmark("input.txt", 1, b) }
func BenchmarkPart2_inputTXT(b *testing.B) { benchmark("input.txt", 2, b) }
func BenchmarkPart1_bigTXT(b *testing.B)   { benchmark("big.txt", 1, b) }
func BenchmarkPart2_bigTXT(b *testing.B)   { benchmark("big.txt", 2, b) }
