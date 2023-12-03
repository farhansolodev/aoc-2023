package main

import (
	"bufio"
	"io"
	"os"
	"regexp"
)

func part2() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	fileinfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	// buffer := make([]byte, fileinfo.Size())

	// _, err = file.Read(buffer)
	// if err != nil && err != io.EOF {
	// 	log.Fatal(err)
	// }

	lookup := map[string]int{
		"one":   1,
		"1":     1,
		"two":   2,
		"2":     2,
		"three": 3,
		"3":     3,
		"four":  4,
		"4":     4,
		"five":  5,
		"5":     5,
		"six":   6,
		"6":     6,
		"seven": 7,
		"7":     7,
		"eight": 8,
		"8":     8,
		"nine":  9,
		"9":     9,
	}

	reader := bufio.NewReaderSize(file, int(fileinfo.Size()))
	compiledRegex := regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine|1|2|3|4|5|6|7|8|9`)
	var sum int = 0

	for {
		line, err := reader.ReadBytes('\n') // this function kinda sucks
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		indicesList := compiledRegex.FindAllIndex(line, -1)
		num1loc := indicesList[0]
		num2loc := indicesList[len(indicesList)-1]
		combinedNumber := lookup[string(line[num1loc[0]:num1loc[1]])]*10 + lookup[string(line[num2loc[0]:num2loc[1]])]
		sum += combinedNumber
	}
	println(sum)
}

// should get 55291
func part2_debug() {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	var sum int = 0 // overflows if smaller than int16
	for {
		line, err := reader.ReadBytes('\n') // byte slices allocated here are small enough to go on stack
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		lastNumberLocs := searchBackwards(line)
		println(string(line[lastNumberLocs[0]:lastNumberLocs[1]]))
		panic("debug baby")
	}
	println(sum)
}

const (
	one   = "one"
	two   = "two"
	three = "three"
	four  = "four"
	five  = "five"
	six   = "six"
	seven = "seven"
	eight = "eight"
	nine  = "nine"
)

/*
todo: change int to int8 and later use int64 for extra indexes
todo: take advantage of each word's spelling to continue early
*/
func searchBackwards(input []byte) (loc []int) { // nolint: gocognit
	right := len(input) - 1
	j1c, j2c, j3c, j4c, j5c, j6c, j7c, j8c, j9c := 0, 0, 0, 0, 0, 0, 0, 0, 0
	j1l, j2l, j3l, j4l, j5l, j6l, j7l, j8l, j9l := -1, -1, -1, -1, -1, -1, -1, -1, -1

	for range input {
		// one
		if input[right] == one[len(one)-1-j1c] {
			// set last as right index if this is a new start for this number
			if j1l == -1 {
				j1l = right
			}
			// return location of word if the counter has reached the end of it
			if j1c == len(one)-1 {
				return []int{right, j1l}
			}
			// increment counter for next run
			j1c++
		} else {
			// idempotently reset counter and last
			j1c = 0
			j1l = -1
		}

		// two
		if input[right] == two[len(two)-1-j2c] {
			if j2l == -1 {
				j2l = right
			}
			if j2c == len(two)-1 {
				return []int{right, j2l}
			}
			j2c++
		} else {
			j2c = 0
			j2l = -1
		}

		// three
		if input[right] == three[len(three)-1-j3c] {
			if j3l == -1 {
				j3l = right
			}
			if j3c == len(three)-1 {
				return []int{right, j3l}
			}
			j3c++
		} else {
			j3c = 0
			j3l = -1
		}

		// four
		if input[right] == four[len(four)-1-j4c] {
			if j4l == -1 {
				j4l = right
			}
			if j4c == len(four)-1 {
				return []int{right, j4l}
			}
			j4c++
		} else {
			j4c = 0
			j4l = -1
		}

		// five
		if input[right] == five[len(five)-1-j5c] {
			if j5l == -1 {
				j5l = right
			}
			if j5c == len(five)-1 {
				return []int{right, j5l}
			}
			j5c++
		} else {
			j5c = 0
			j5l = -1
		}

		// six
		if input[right] == six[len(six)-1-j6c] {
			if j6l == -1 {
				j6l = right
			}
			if j6c == len(six)-1 {
				return []int{right, j6l}
			}
			j6c++
		} else {
			j6c = 0
			j6l = -1
		}

		// seven
		if input[right] == seven[len(seven)-1-j7c] {
			if j7l == -1 {
				j7l = right
			}
			if j7c == len(seven)-1 {
				return []int{right, j7l}
			}
			j7c++
		} else {
			j7c = 0
			j7l = -1
		}

		// eight
		if input[right] == eight[len(eight)-1-j8c] {
			if j8l == -1 {
				j8l = right
			}
			if j8c == len(eight)-1 {
				return []int{right, j8l}
			}
			j8c++
		} else {
			j8c = 0
			j8l = -1
		}

		// nine
		if input[right] == nine[len(nine)-1-j9c] {
			if j9l == -1 {
				j9l = right
			}
			if j9c == len(nine)-1 {
				return []int{right, j9l}
			}
			j9c++
		} else {
			j9c = 0
			j9l = -1
		}

		right--
	}
	return nil
}
