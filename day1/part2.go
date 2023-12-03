package main

import (
	"bufio"
	"io"
	"os"
)

func part2(file *os.File) (sum int) {
	defer file.Close()
	reader := bufio.NewReader(file)

	var line []byte
	var err error
	for {
		line, err = reader.ReadBytes('\n') // byte slices allocated here are small enough to go on stack so no gc pressure
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		nums := search(line)
		sum += int(nums[0])*10 + int(nums[1])
	}
	return sum
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
todo: take advantage of each word's spelling to continue early
*/
func search(line []byte) (nums [2]int8) { // nolint: gocognit
	// none of these numbers will be >= 127, so int8 won't overflow
	var j1c, j2c, j3c, j4c, j5c, j6c, j7c, j8c, j9c int8 = 0, 0, 0, 0, 0, 0, 0, 0, 0
	var j1l, j2l, j3l, j4l, j5l, j6l, j7l, j8l, j9l int8 = -1, -1, -1, -1, -1, -1, -1, -1, -1
	var i1c, i2c, i3c, i4c, i5c, i6c, i7c, i8c, i9c int8 = 0, 0, 0, 0, 0, 0, 0, 0, 0
	var i1l, i2l, i3l, i4l, i5l, i6l, i7l, i8l, i9l int8 = -1, -1, -1, -1, -1, -1, -1, -1, -1
	// 36 bytes on stack^

	nums = [2]int8{-1, -1}

	right := int8(len(line) - 1)
	for left := int8(0); left < int8(len(line)); left++ {
		// break if both numbers have been found
		if nums[0] != -1 && nums[1] != -1 {
			break
		}

		// skip right if its already been found
		if nums[1] != -1 {
			goto left
		}

		// found a digit
		if nums[1] == -1 && line[right] >= '1' && line[right] <= '9' {
			nums[1] = int8(line[right] - '0')
			goto left
		}

		// one
		if j1c != -1 {
			switch line[right] {
			case one[int8(len(one))-1-j1c]:
				// set last as right index if this is a new start for this number
				if j1l == -1 {
					j1l = right
				}
				// set done flag and found number if the counter has reached the end of it, then go for next run
				if j1c == int8(len(one))-1 {
					nums[1] = 1
					j1c = -1
					goto left
				}
				// increment counter for next run
				j1c++
			case one[int8(len(one))-1]: // every character could be the start of a new run
				j1c = 1
				j1l = right
			default:
				// idempotently reset counter and last
				j1c = 0
				j1l = -1
			}
		}

		// two
		if j2c != -1 {
			switch line[right] {
			case two[int8(len(two))-1-j2c]:
				if j2l == -1 {
					j2l = right
				}
				if j2c == int8(len(two))-1 {
					nums[1] = 2
					j2c = -1
					goto left
				}
				j2c++
			case two[int8(len(two))-1]:
				j2c = 1
				j2l = right
			default:
				j2c = 0
				j2l = -1
			}
		}

		// three
		if j3c != -1 {
			switch line[right] {
			case three[int8(len(three))-1-j3c]:
				if j3l == -1 {
					j3l = right
				}
				if j3c == int8(len(three))-1 {
					nums[1] = 3
					j3c = -1
					goto left
				}
				j3c++
			case three[int8(len(three))-1]:
				j3c = 1
				j3l = right
			default:
				j3c = 0
				j3l = -1
			}
		}

		// four
		if j4c != -1 {
			switch line[right] {
			case four[int8(len(four))-1-j4c]:
				if j4l == -1 {
					j4l = right
				}
				if j4c == int8(len(four))-1 {
					nums[1] = 4
					j4c = -1
					goto left
				}
				j4c++
			case four[int8(len(four))-1]:
				j4c = 1
				j4l = right
			default:
				j4c = 0
				j4l = -1
			}
		}

		// five
		if j5c != -1 {
			switch line[right] {
			case five[int8(len(five))-1-j5c]:
				if j5l == -1 {
					j5l = right
				}
				if j5c == int8(len(five))-1 {
					nums[1] = 5
					j5c = -1
					goto left
				}
				j5c++
			case five[int8(len(five))-1]:
				j5c = 1
				j5l = right
			default:
				j5c = 0
				j5l = -1
			}
		}

		// six
		if j6c != -1 {
			switch line[right] {
			case six[int8(len(six))-1-j6c]:
				if j6l == -1 {
					j6l = right
				}
				if j6c == int8(len(six))-1 {
					nums[1] = 6
					j6c = -1
					goto left
				}
				j6c++
			case six[int8(len(six))-1]:
				j6c = 1
				j6l = right
			default:
				j6c = 0
				j6l = -1
			}
		}

		// seven
		if j7c != -1 {
			switch line[right] {
			case seven[int8(len(seven))-1-j7c]:
				if j7l == -1 {
					j7l = right
				}
				if j7c == int8(len(seven))-1 {
					nums[1] = 7
					j7c = -1
					goto left
				}
				j7c++
			case seven[int8(len(seven))-1]:
				j7c = 1
				j7l = right
			default:
				j7c = 0
				j7l = -1
			}
		}

		// eight
		if j8c != -1 {
			switch line[right] {
			case eight[int8(len(eight))-1-j8c]:
				if j8l == -1 {
					j8l = right
				}
				if j8c == int8(len(eight))-1 {
					nums[1] = 8
					j8c = -1
					goto left
				}
				j8c++
			case eight[int8(len(eight))-1]:
				j8c = 1
				j8l = right
			default:
				j8c = 0
				j8l = -1
			}
		}

		// nine
		if j9c != -1 {
			switch line[right] {
			case nine[int8(len(nine))-1-j9c]:
				if j9l == -1 {
					j9l = right
				}
				if j9c == int8(len(nine))-1 {
					nums[1] = 9
					j9c = -1
					goto left
				}
				j9c++
			case nine[int8(len(nine))-1]:
				j9c = 1
				j9l = right
			default:
				j9c = 0
				j9l = -1
			}
		}

	left:
		// optimistically break if both numbers have been found
		if nums[0] != -1 && nums[1] != -1 {
			break
		}

		// skip to next run if left has already been found
		if nums[0] != -1 {
			right--
			continue
		}

		// found a digit
		if nums[0] == -1 && line[left] >= '1' && line[left] <= '9' {
			nums[0] = int8(line[left] - '0')
			right--
			continue
		}

		// one
		if i1c != -1 { // only run if number not yet found
			switch line[left] {
			case one[i1c]:
				// set last as left index if this is a new start for this number
				if i1l == -1 {
					i1l = left
				}
				// set done flag and location of word if the counter has reached the end of it, then go for next run
				if i1c == int8(len(one))-1 {
					nums[0] = 1
					i1c = -1
					right--
					continue
				}
				// increment counter for next run
				i1c++
			case one[0]: // every character could be the start of a new run
				i1c = 1
				i1l = left
			default:
				// idempotently reset counter and last
				i1c = 0
				i1l = -1
			}
		}

		// two
		if i2c != -1 {
			switch line[left] {
			case two[i2c]:
				if i2l == -1 {
					i2l = left
				}
				if i2c == int8(len(two))-1 {
					nums[0] = 2
					i2c = -1
					right--
					continue
				}
				i2c++
			case two[0]:
				i2c = 1
				i2l = left
			default:
				i2c = 0
				i2l = -1
			}
		}

		// three
		if i3c != -1 {
			switch line[left] {
			case three[i3c]:
				if i3l == -1 {
					i3l = left
				}
				if i3c == int8(len(three))-1 {
					nums[0] = 3
					i3c = -1
					right--
					continue
				}
				i3c++
			case three[0]:
				i3c = 1
				i3l = left
			default:
				i3c = 0
				i3l = -1
			}
		}

		// four
		if i4c != -1 {
			switch line[left] {
			case four[i4c]:
				if i4l == -1 {
					i4l = left
				}
				if i4c == int8(len(four))-1 {
					nums[0] = 4
					i4c = -1
					right--
					continue
				}
				i4c++
			case four[0]:
				i4c = 1
				i4l = left
			default:
				i4c = 0
				i4l = -1
			}
		}

		// five
		if i5c != -1 {
			switch line[left] {
			case five[i5c]:
				if i5l == -1 {
					i5l = left
				}
				if i5c == int8(len(five))-1 {
					nums[0] = 5
					i5c = -1
					right--
					continue
				}
				i5c++
			case five[0]:
				i5c = 1
				i5l = left
			default:
				i5c = 0
				i5l = -1
			}
		}

		// six
		if i6c != -1 {
			switch line[left] {
			case six[i6c]:
				if i6l == -1 {
					i6l = left
				}
				if i6c == int8(len(six))-1 {
					nums[0] = 6
					i6c = -1
					right--
					continue
				}
				i6c++
			case six[0]:
				i6c = 1
				i6l = left
			default:
				i6c = 0
				i6l = -1
			}
		}

		// seven
		if i7c != -1 {
			switch line[left] {
			case seven[i7c]:
				if i7l == -1 {
					i7l = left
				}
				if i7c == int8(len(seven))-1 {
					nums[0] = 7
					i7c = -1
					right--
					continue
				}
				i7c++
			case seven[0]:
				i7c = 1
				i7l = left
			default:
				i7c = 0
				i7l = -1
			}
		}

		// eight
		if i8c != -1 {
			switch line[left] {
			case eight[i8c]:
				if i8l == -1 {
					i8l = left
				}
				if i8c == int8(len(eight))-1 {
					nums[0] = 8
					i8c = -1
					right--
					continue
				}
				i8c++
			case eight[0]:
				i8c = 1
				i8l = left
			default:
				i8c = 0
				i8l = -1
			}
		}

		// nine
		if i9c != -1 {
			switch line[left] {
			case nine[i9c]:
				if i9l == -1 {
					i9l = left
				}
				if i9c == int8(len(nine))-1 {
					nums[0] = 9
					i9c = -1
					right--
					continue
				}
				i9c++
			case nine[0]:
				i9c = 1
				i9l = left
			default:
				i9c = 0
				i9l = -1
			}
		}

		right--
	}

	return nums
}
