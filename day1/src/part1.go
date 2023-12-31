package day1

import (
	"bufio"
	"io"
	"os"
)

func Part1(file *os.File) (sum int) {
	reader := bufio.NewReader(file)
	tmp := []byte{0, 0}

	for {
		line, err := reader.ReadBytes('\n') // byte slices allocated here are small enough to go on stack so no gc pressure
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		rightMarker := len(line) - 1
		for leftMarker := range line {
			if tmp[0] != 0 && tmp[1] != 0 {
				break
			}

			leftByte := line[leftMarker]
			rightByte := line[rightMarker]

			if tmp[0] == 0 && leftByte >= '0' && leftByte <= '9' {
				tmp[0] = leftByte
			}
			if tmp[1] == 0 && rightByte >= '0' && rightByte <= '9' {
				tmp[1] = rightByte
			}

			rightMarker--
		}
		num := (tmp[0]-'0')*10 + tmp[1] - '0'
		tmp[0] = 0
		tmp[1] = 0
		sum += int(num)
	}

	return sum
}
