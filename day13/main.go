package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	readFile, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	boards := strings.Split(string(readFile), "\n\n")

	sum := 0

	for _, board := range boards {

		parsed := strings.Split(board, "\n")

		sum += checkVerticalMirror(parsed)
		sum += checkHorizontalMirror(parsed)
	}

	fmt.Println(sum)
}

func checkHorizontalMirror(board []string) int {
	for i := 0; i < len(board)-1; i++ {
		if checkMirrorRows(board, i, 0) {

			r := 0
			if len(board)-2-i > i {
				r = i
			} else {
				r = len(board) - 2 - i
			}

			mirror := true
			for j := 1; j <= r; j++ {
				if !checkMirrorRows(board, i, j) {
					mirror = false
				}
			}

			if mirror {
				return (i + 1) * 100
			}
		}
	}

	return 0
}

func checkVerticalMirror(board []string) int {
	for i := 0; i < len(board[0])-1; i++ {
		if checkMirrorColumns(board, i, 0) {

			r := 0
			if len(board[0])-2-i > i {
				r = i
			} else {
				r = len(board[0]) - 2 - i
			}

			mirror := true
			for j := 1; j <= r; j++ {
				if !checkMirrorColumns(board, i, j) {
					mirror = false
				}
			}

			if mirror {
				return i + 1
			}
		}
	}

	return 0
}

func checkMirrorColumns(board []string, columnNumber int, offset int) bool {
	for i := 0; i < len(board); i++ {
		if board[i][columnNumber-offset] != board[i][columnNumber+1+offset] {
			return false
		}
	}
	return true
}

func checkMirrorRows(board []string, rowNumber int, offset int) bool {
	return board[rowNumber-offset] == board[rowNumber+1+offset]
}
