package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var board []string

func main() {
	readFile, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	board = strings.Split(string(readFile), "\n")
	sum := 0

	for i := 0; i < len(board); i++ {
		wasDigit := false
		startIndex := 0

		for j := 0; j < len(board[i]); j++ {

			r := rune(board[i][j])

			if unicode.IsDigit(r) {
				if !wasDigit {
					startIndex = j
				}

				wasDigit = true

				if j == len(board[i])-1 {
					if CheckIfPart(i, j+1, j-startIndex+1) {
						val, _ := strconv.Atoi(board[i][startIndex : j+1])

						sum += val
					}
				}
			} else {
				if wasDigit {
					if CheckIfPart(i, j, j-startIndex) {

						val, _ := strconv.Atoi(board[i][startIndex:j])
						sum += val
					}
					wasDigit = false
				}
			}

		}
	}
	fmt.Println(sum)
}

func CheckIfPart(x, y, length int) bool {
	bounds_x, bounds_y := x-1, y-length-1

	for i := bounds_x; i < bounds_x+3; i++ {
		for j := bounds_y; j < bounds_y+length+2; j++ {

			if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) {
				continue
			}

			if IsSymbol(rune(board[i][j])) {
				return true
			}
		}
	}

	return false
}

func IsSymbol(r rune) bool {
	return r != rune(46) && !unicode.IsDigit(r)
}
