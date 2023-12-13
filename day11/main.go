package main

import (
	"fmt"
	"os"
	"strings"
)

var board []string

type galaxy struct {
	x int
	y int
}

func main() {
	readFile, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	board = strings.Split(string(readFile), "\n")

	galaxies := []galaxy{}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '#' {
				galaxies = append(galaxies, galaxy{x: j, y: i})
			}
		}
	}

	fmt.Println(galaxies)

	doubleColumns := []int{}
	doubleRows := []int{}

	for i := 0; i < len(board); i++ {
		if board[i] == emptyRow(len(board[i])) {
			doubleRows = append(doubleRows, i)
		}
	}

	fmt.Println(doubleRows)

	for i := 0; i < len(board[1]); i++ {
		empty := true
		for j := 0; j < len(board); j++ {
			// fmt.Println(j, i, string(board[j][i]))
			// fmt.Scanln()

			if board[j][i] == '#' {
				empty = false
				break
			}
		}

		if empty {
			doubleColumns = append(doubleColumns, i)
		}
	}

	fmt.Println(doubleColumns)
}

func emptyRow(length int) (row string) {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteRune('.')
	}

	return sb.String()
}
